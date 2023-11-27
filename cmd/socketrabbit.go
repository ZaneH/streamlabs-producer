package main

import (
	"log"
	"socketrabbit/internal/rabbit"
	"socketrabbit/internal/third-party/streamlabs"
	"time"
)

func main() {
	log.Printf("Starting...")

	// initialize RabbitMQ client
	rc := rabbit.NewClient()
	err := rc.Setup()
	defer rc.Close()

	if err != nil {
		log.Fatalf("Error creating RabbitMQ client: %s", err)
		panic(err)
	}

	// 1 exchange per 3rd party service
	err = rc.NewExchange("server.streamlabs")
	if err != nil {
		log.Fatalf("Error declaring RabbitMQ exchange: %s", err)
		panic(err)
	}

	// 1 queue per user
	_, err = rc.NewQueue("frontend.consumer.1")
	if err != nil {
		log.Fatalf("Error declaring RabbitMQ queue: %s", err)
		panic(err)
	}

	// bind queue + exchange with key "1" for demo purposes
	// key 1 == user 1
	err = rc.Channel.QueueBind("frontend.consumer.1", "1", "server.streamlabs", false, nil)
	if err != nil {
		log.Fatalf("Error binding RabbitMQ queue: %s", err)
		panic(err)
	}

	c := streamlabs.NewClient()
	err = c.Setup(rc)
	defer c.Close()

	if err != nil {
		log.Fatalf("Error connecting to Streamlabs: %s", err)
	}

	time.Sleep(10 * time.Second)
}
