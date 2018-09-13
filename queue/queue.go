package queue

import (
	"reflect"
	"time"
)

type Request struct {
	Channel   chan int
	Url       string
	Score     int
	StartTime time.Time
}

type RequestLoad struct {
	url            string
	processingTime int
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
			q.removeActiveItemByIndex(i)
		}
	}
}

func (q *Queue) UpdateRequestQueue(addRequestChannel chan *Request, removeRequestChannel chan *Request) {
	cases := make([]reflect.SelectCase, 2)
	cases[0] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(addRequestChannel)}
	cases[1] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(removeRequestChannel)}
	for {
		chosen, value, _ := reflect.Select(cases)
		request := value.Interface().(*Request)
		if chosen == 0 {
			q.priorityQueue = append(q.priorityQueue, request)
		} else if chosen == 1 {
			q.removeActiveItem(request)
		}

	}
}

func (q *Queue) HandleQueue() {
	for {
		if availableSlots() > 0 {
			q.activateRequests()
		}

	}
}
func (q *Queue) activateRequests() {
	for _, request := range q.priorityQueue {
		if availableSlots() > 0 && request.Score == 0 {
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
