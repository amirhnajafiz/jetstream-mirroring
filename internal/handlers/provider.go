package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func (h Handler) Provider(host string) error {
	// connect to NATS server for publishing
	nc, err := nats.Connect(host)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS cluster: %w", err)
	}

	// creating a JetStream connection
	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("failed to open jetstream connection: %w", err)
	}

	// publish message
	for {
		_, err = js.Publish(h.Stream.SubjectName, []byte("input buffer for testing message"))
		if err != nil {
			log.Println(fmt.Errorf("failed to publish on %s: %w", host, err))
		}

		// sleep for interval
		time.Sleep(h.ProviderInterval)
	}
}
