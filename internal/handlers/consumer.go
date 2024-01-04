package handlers

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func (h Handler) Consumer(host string) error {
	// connect to NATS server for subscribe
	nc, err := nats.Connect(host)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS cluster: %w", err)
	}

	// creating a JetStream connection
	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("failed to open jetstream connection: %w", err)
	}

	// subscribing on subject
	_, _ = js.Subscribe(h.Stream.SubjectName, func(msg *nats.Msg) {
		// send acknowledgement
		if er := msg.Ack(); er != nil {
			log.Println(fmt.Errorf("failed to send ack to %s: %w", host, err))

			return
		}

		log.Println(fmt.Sprintf("consumed %d on %s", len(msg.Data), host))
	})
}
