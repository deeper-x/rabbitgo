package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Engine is the rabbit engine
type Engine struct {
	Connection *amqp.Connection
}

// UserInput definitions
type UserInput struct {
	action *string
	msg    *string
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
