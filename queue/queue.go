package queue

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Request struct {
	Channel   chan int
	Url       string
	Ip        string
	Score     int64
	StartTime time.Time
}

type RequestLoad struct {
	processingTime int64
	count          int
}

type Queue struct {
	priorityQueue []*Request
	requestLoad   map[string]*RequestLoad
	FromIpScore   map[string]int64
}

var avgLoadTime int64 = 0

const maxRequests = 20

var activeQueue []*Request

func totalApproxAvgUpdate(value int64) {
	var samplecount int64 = 50000
	avgLoadTime = avgLoadTime*(samplecount-1)/samplecount + value
}

func (q *Queue) removeQueueItemByIndex(i int) {

	q.priorityQueue = q.priorityQueue[:i+copy(q.priorityQueue[i:], q.priorityQueue[i+1:])]
}

func (q *Queue) removeActiveItemByIndex(i int) {

	activeQueue = activeQueue[:i+copy(activeQueue[i:], activeQueue[i+1:])]
}

func (q *Queue) removeQueueItem(request *Request) {

	for i, queuedRequest := range q.priorityQueue {
		if queuedRequest == request {
			q.removeQueueItemByIndex(i)
		}
	}
}

func (q *Queue) removeActiveItem(request *Request) {

	for i, queuedRequest := range activeQueue {
		if queuedRequest == request {
			q.updateLoad(request)
			q.removeActiveItemByIndex(i)
			spew.Dump(q)
		}
	}
}

func (q *Queue) UpdateRequestQueue(addRequestChannel chan *Request, removeRequestChannel chan *Request) {

	timer := time.NewTimer(1 * time.Second)

	for {
		select {
		//A new connection is accepted
		case request := <-addRequestChannel:
			q.addRequest(request)

		//A connection is done
		case request := <-removeRequestChannel:
			q.removeActiveItem(request)

		//The stats timer has ran out
		case <-timer.C:
			q.sendStats()
			timer = time.NewTimer(1 * time.Second)
		default:
		}

		q.activateRequests()
	}
}

func (q *Queue) sendStats() {
	/*type stat struct {
		ipaddress         string `json:"ip_addres"`
		timeSpentInQueue  int64
		requestsPerSecond int64
	}

	stats := make([]stat, 0, len(q.priorityQueue))
	//	q.priorityQueue*/
}

func (q *Queue) addRequest(request *Request) {
	q.priorityQueue = append(q.priorityQueue, request)
	q.updateFromIpScore(request)
	request.Score = q.getScore(request)
	// todo: adjust current score
}

func (q *Queue) getIpScore(ip string) int64 {
	return q.FromIpScore[ip]
}
func (q *Queue) getLoadTime(url string) int64 {
	if q.requestLoad[url] != nil {
		return q.requestLoad[url].processingTime
	}
	return int64(0)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (q *Queue) getScore(request *Request) int64 {
	precision := int64(100)
	score := q.getIpScore(request.Ip)
	loadTime := q.getLoadTime(request.Url)
	if loadTime > 0 {
		tLoad := loadTime * precision
		tavg := avgLoadTime * precision
		diff := tLoad / tavg
		score = (score * precision * max(diff, precision)) / 100
	}
	return score
}

func (q *Queue) updateLoad(request *Request) {
	timediff := time.Now().Sub(request.StartTime).Nanoseconds()
	totalApproxAvgUpdate(timediff)
	if q.requestLoad == nil {
		q.requestLoad = map[string]*RequestLoad{}
	}

	if _, exist := q.requestLoad[request.Url]; exist {
		q.requestLoad[request.Url].processingTime = ((q.requestLoad[request.Url].processingTime * int64(q.requestLoad[request.Url].count)) + timediff) / (int64(q.requestLoad[request.Url].count + 1))
		q.requestLoad[request.Url].count = q.requestLoad[request.Url].count + 1
	}

	requestLoad := RequestLoad{}
	requestLoad.processingTime = time.Now().Sub(request.StartTime).Nanoseconds()
	requestLoad.count = 1
	q.requestLoad[request.Url] = &requestLoad
}

func (q *Queue) updateFromIpScore(request *Request) {

	if q.FromIpScore == nil {
		q.FromIpScore = map[string]int64{}
	}

	q.FromIpScore[request.Ip] = q.FromIpScore[request.Ip] + 1

}

func (q *Queue) activateRequests() {
	if availableSlots() > 0 {
		for _, request := range q.priorityQueue {
			if request.Score == 0 {

				request.StartTime = time.Now()
				q.removeQueueItem(request)

				activeQueue = append(activeQueue, request)
				request.Channel <- 1

			}
		}
	}
	q.decrementScores()
}

func (q *Queue) decrementScores() {

	for i, request := range q.priorityQueue {
		if request.Score > 0 {
			q.priorityQueue[i].Score = request.Score - 1
		}
	}
}

func availableSlots() int {
	slots := maxRequests - len(activeQueue)
	return slots
}
