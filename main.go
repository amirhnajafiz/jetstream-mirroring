package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

const (
	nats1 = "nats://0.0.0.0:6222"
	nats2 = "nats://0.0.0.0:6223"
)

func main() {
	{
		// Connect to NATS server 1
		nc, _ := nats.Connect(nats1)
		_, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}
	}
	{
		// Connect to NATS server 1
		nc, _ := nats.Connect(nats2)
		_, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}
	}
}
