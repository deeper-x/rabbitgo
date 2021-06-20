package main

import (
	"testing"
)

func TestStart(t *testing.T) {
	e := Engine{}
	err := e.Start()
	if err != nil {
		t.Error(err)
	}

	defer e.Connection.Close()
}

func TestSend(t *testing.T) {
	e := Engine{}
	err := e.Start()
	if err != nil {
		t.Error(err)
	}
	defer e.Connection.Close()

	err = e.send("demo")

	if err != nil {
		t.Error(err)
	}
}

func TestPub(t *testing.T) {
	e := Engine{}
	err := e.Start()
	if err != nil {
		t.Error(err)
	}
	defer e.Connection.Close()

	err = e.Pub("demo")

	if err != nil {
		t.Error(err)
	}
}
