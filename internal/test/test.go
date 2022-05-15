package main

import (
	"log"
	"time"

	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/nats-io/nats.go"
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

	ticker := time.NewTimer(1 * time.Second)
	for i := range ticker.C {
		_, err = js.Publish(cfg.SubjectName, []byte(message))
		if err != nil {
			log.Printf("[Test %d] Error: %s\n", i.Second(), err.Error())
		} else {
			log.Printf("[Test %d] Done\n", i.Second())
		}
	}
}
