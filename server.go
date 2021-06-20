package main

import (
	"log"
)

// Sub is the identity management provider
func (e Engine) Sub() {
	ch, err := e.Connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("auth_queue", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for m := range msgs {
			log.Printf("#TODO - Authenticating identity: %s", m.Body)
		}
	}()

	sbscr := make(chan bool)

	<-sbscr
}
