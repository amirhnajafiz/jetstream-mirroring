package handlers

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// Bootstrap method sets up the NATS clusters for testing
func (h Handler) Bootstrap(hosts []string) []error {
	errors := make([]error, 0)

	for _, host := range hosts {
		// connect to NATS server
		nc, err := nats.Connect(host)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to connect to %s cluster", host))

			continue
		}

		// creating a JetStream connection
		js, err := nc.JetStream()
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to create jetstream connection in %s cluster", host))

			continue
		}

		// create a stream instance
		err = h.createStream(js)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to create stream instance in %s cluster", host))

			continue
		}

		log.Println(fmt.Sprintf("created stream %q and subjects %q on %s cluster", "", "", host))
	}

	return errors
}

// this method creates our stream in NATS
func (h Handler) createStream(js nats.JetStreamContext) error {
	stream, err := js.StreamInfo("")
	if err != nil {
		return err
	}

	if stream == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     "",
			Subjects: []string{""},
		})
		if err != nil {
			return err
		}
	}

	return nil
}
