package main

import (
	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/nats-io/nats.go"
	"log"
)

const (
	// message to publish
	message = "snapp.cab"
)

func main() {
	cfg := config.Load()

	// Connect to NATS server
	nc, err := nats.Connect(cfg.Nat1)
	if err != nil {
		log.Fatal(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < 11; i++ {
		_, err = js.Publish(cfg.SubjectName, []byte(message))
		if err != nil {
			log.Printf("[Test %d] Error: %s\n", i, err.Error())
		} else {
			log.Printf("[Test %d] Done\n", i)
		}
	}
}
