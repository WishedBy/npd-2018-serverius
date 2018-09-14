package queue

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Request struct {
	Channel   chan int
	Url       string
	Score     int
	StartTime time.Time
}

type RequestLoad struct {
	url            string
	processingTime int64
	count          int
}
type Queue struct {
	priorityQueue []*Request
	requestLoad   []*RequestLoad
}

const maxRequests = 20

var activeQueue []*Request

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
	for {
		select {
		//A new connection is accepted
		case request := <-addRequestChannel:
			q.priorityQueue = append(q.priorityQueue, request)

		//A connection is done
		case request := <-removeRequestChannel:
			q.removeActiveItem(request)
		default:
		}

		q.handleQueue()
	}
}

func (q *Queue) updateLoad(request *Request) {
	timediff := time.Now().Sub(request.StartTime).Nanoseconds()
	for _, requestLoad := range q.requestLoad {
		if requestLoad.url == request.Url {
			requestLoad.processingTime = ((requestLoad.processingTime * int64(requestLoad.count)) + timediff) / (int64(requestLoad.count + 1))
			requestLoad.count = requestLoad.count + 1
			return
		}

	}
	requestLoad := RequestLoad{}
	requestLoad.url = request.Url
	requestLoad.processingTime = time.Now().Sub(request.StartTime).Nanoseconds()
	requestLoad.count = 1
	q.requestLoad = append(q.requestLoad, &requestLoad)
}

func (q *Queue) handleQueue() {
	if availableSlots() > 0 {
		q.activateRequests()
	}
}
func (q *Queue) activateRequests() {
	for _, request := range q.priorityQueue {
		if request.Score == 0 {
			request.StartTime = time.Now()
			q.removeQueueItem(request)
			activeQueue = append(activeQueue, request)
			request.Channel <- 1

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
