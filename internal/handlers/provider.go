package handlers

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

func (h Handler) Provider(host string) error {
	intervalCh := make(chan int)
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

	// publish message
	for {
		_, err = js.Publish(h.Stream.SubjectName, []byte("input buffer for testing message"))
		if err != nil {
			log.Println(fmt.Errorf("failed to publish on %s: %w", host, err))
		}

		go func() {
			// sleep for interval
			time.Sleep(h.ProviderInterval)
			intervalCh <- 0
		}()

		select {
		case <-intervalCh:
			continue
		case <-signalCh:
			os.Exit(0)
		}
	}
}
