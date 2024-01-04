package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func (h Handler) Provider(host string) error {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

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

	mgs := []byte("input buffer for testing message")

	for {
		// create a new context
		ctx, cancel := context.WithTimeout(context.Background(), h.ProviderInterval)

		// publish
		_, err = js.Publish(h.Stream.SubjectName, mgs)
		if err != nil {
			log.Println(fmt.Errorf("failed to publish on %s: %w", host, err))
		}

		log.Println(fmt.Sprintf("published %d bytes on %s", len(mgs), host))

		select {
		case <-ctx.Done():
			cancel()
		case <-signalCh:
			os.Exit(0)
		}
	}
}
