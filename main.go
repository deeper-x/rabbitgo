package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	ia := inputArgs()
	e := Engine{}

	err := e.Start()
	if err != nil {
		log.Fatal(err)
	}

	defer e.Connection.Close()

	switch *ia.action {

	case "pub":
		e.Pub(*ia.msg)

	case "sub":
		e.Sub()

	default:
		log.Println("no accepted action...")
		os.Exit(1)
	}

}

func inputArgs() *UserInput {
	action := flag.String("action", "sub", "sub for subscribing, pub for publishing")
	msg := flag.String("msg", "no msg", "input message to publish")

	flag.Parse()

	return &UserInput{
		action: action,
		msg:    msg,
	}
}
