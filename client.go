package main

import (
	"log"

	"github.com/streadway/amqp"
)

// send message to broker
func (e Engine) send(msg string) error {
	ch, err := e.Connection.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("auth_queue", false, false, false, false, nil)
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

// Pub is the message carrier
func (e Engine) Pub(msg string) error {
	err := e.send(msg)
	if err != nil {
		return err
	}

	return nil
}
