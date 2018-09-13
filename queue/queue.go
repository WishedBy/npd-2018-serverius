package queue

type Routine struct {
	channel *chan string
	score   int
}
type Queue struct {
	priorityQueue []Routine
}

func (q *Queue) nextRequest() {

}
func (q *Queue) setRequest(channel chan string, score int) {

}
