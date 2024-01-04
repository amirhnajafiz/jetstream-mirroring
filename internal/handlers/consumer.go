package handlers

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func (h Handler) Consumer(host string) error {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

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

		log.Println(fmt.Sprintf("consumed %d bytes from %s", len(msg.Data), host))
	})

	// wait for cancel signal
	<-signalCh
	os.Exit(0)

	return nil
}
