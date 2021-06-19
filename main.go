package main

import (
	"flag"
	"os"
)

func main() {
	action := flag.String("action", "send", "send for subscribing, receive for publishing")
	msg := flag.String("msg", "no msg", "input message to publish")

	flag.Parse()

	switch *action {
	case "send":
		Pub(*msg)

	case "receive":
		Sub()

	default:
		os.Exit(1)
	}

}
