package main

import (
	"log"
	"math/rand"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

const (
	address   = "127.0.0.1:7233"
	namespace = "default"
	taskQueue = "workshop"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create client
	c, err := client.Dial(client.Options{
		HostPort:  address,
		Namespace: namespace,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// Create worker
	w := worker.New(c, taskQueue, worker.Options{})

	// Register workflows and activities
	register(w)

	// Start worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
