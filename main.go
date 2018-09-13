package main

import (
	"log"
	"npd/prototype"
	"npd/queue"
)

func main() {
	q := queue.Queue{}
	addRequestChannel := make(chan *queue.Request)
	removeRequestChannel := make(chan *queue.Request)

	go q.UpdateRequestQueue(addRequestChannel, removeRequestChannel)
	go q.HandleQueue()

	npd := &prototype.Npd{}
	npd.SetAddQueueChannel(addRequestChannel)
	npd.SetRemoveQueueChannel(removeRequestChannel)
	log.Println("Starting server")

	log.Fatal(npd.Start())
}
