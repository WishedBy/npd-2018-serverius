package main

import (
	"log"
	"npd/prototype"
	"npd/queue"
)

func main() {
	q := queue.Queue{}
	addRequestChannel := make(chan queue.Request)

	go q.AddRequest(addRequestChannel)
	go q.HandleQueue()

	npd := &prototype.Npd{}
	npd.SetAddQueueChannel(addRequestChannel)
	log.Println("Starting server")

	log.Fatal(npd.Start())
}
