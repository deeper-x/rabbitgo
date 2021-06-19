package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Pub is the publisher
func Pub(msg string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println(err)
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("queue_declared", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
