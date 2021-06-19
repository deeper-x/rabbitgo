package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Sub is the subscriber
func Sub() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("queue_declared", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	rcvr := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("%s", d.Body)
		}
	}()

	<-rcvr
}
