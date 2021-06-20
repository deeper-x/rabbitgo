package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	action := flag.String("action", "sub", "sub for subscribing, pub for publishing")
	msg := flag.String("msg", "no msg", "input message to publish")

	flag.Parse()

	e := Engine{}

	err := e.Start()
	if err != nil {
		log.Fatal(err)
	}

	defer e.Connection.Close()

	switch *action {

	case "pub":
		log.Println("Sending identity")
		e.Pub(*msg)

	case "sub":
		log.Println("Listening...")
		e.Sub()

	default:
		log.Println("no accepted action...")
		os.Exit(1)
	}

}
