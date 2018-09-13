package queue

import (
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

var maxRequests = 20
var currentRequests = 0

func (q *Queue) removeQueueItemByIndex(i int) {

	q.priorityQueue = q.priorityQueue[:i+copy(q.priorityQueue[i:], q.priorityQueue[i+1:])]
}
func (q *Queue) removeQueueItem(request *Request) {

	for i, queuedRequest := range q.priorityQueue {
		if queuedRequest == request {
			q.removeQueueItemByIndex(i)
		}
	}
}

func (q *Queue) AddRequest(addRequestChannel chan *Request) {
	for {
		request := <-addRequestChannel
		q.priorityQueue = append(q.priorityQueue, request)

	}
}
func (q *Queue) RemoveRequest(removeRequestChannel chan *Request) {
	for {
		request := <-removeRequestChannel
		q.removeQueueItem(request)
		decrementCounter()
	}
}

func (q *Queue) HandleQueue() {
	for {
		if availableSlots() > 0 {
			q.allowRequests()
		}

	}
}
func (q *Queue) allowRequests() {
	for _, request := range q.priorityQueue {
		if availableSlots() > 0 && request.Score == 0 && request.StartTime.IsZero() {
			incrementCounter()
			request.StartTime = time.Now()
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
	slots := maxRequests - currentRequests
	return slots
}

func incrementCounter() int {
	currentRequests = currentRequests + 1
	return currentRequests
}

func decrementCounter() int {
	currentRequests = currentRequests - 1
	return currentRequests
}
