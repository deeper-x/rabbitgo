package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Engine is the rabbit engine
type Engine struct {
	Connection *amqp.Connection
}

// Send message to broker
func (e Engine) Send(msg string) error {
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

// Start configure instance connection
func (e *Engine) Start() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
		return err
	}

	e.Connection = conn
	return nil
}
