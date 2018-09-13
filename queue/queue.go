package queue

import "log"

type Request struct {
	Channel chan int
	Url     string
	Score   int
}

type RequestLoad struct {
	url            string
	processingTime int
}
type Queue struct {
	priorityQueue []Request
	requestLoad   []RequestLoad
}

var maxRequests = 20
var currentRequests = 0

func (q *Queue) nextRequest() {

}
func (q *Queue) setRequest(channel chan int, score int) {

}

func (q *Queue) AddRequest(addRequestChannel chan Request) {
	for {
		request := <-addRequestChannel
		q.priorityQueue = append(q.priorityQueue, request)

		log.Println(request.Url)
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
		if availableSlots() > 0 && request.Score == 0 {
			incrementCounter()
			request.Channel <- 1
		}
	}
	q.decrementScores()
}
func (q *Queue) decrementScores() {

	for i, request := range q.priorityQueue {
		if request.Score > 0 {
			q.priorityQueue[i].Score = request.Score - 1

			log.Println(q.priorityQueue[i].Score)
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

func decrementCounter(addRequestChannel chan Request) int {
	currentRequests = currentRequests - 1
	return currentRequests
}

func requestDone() {

}
