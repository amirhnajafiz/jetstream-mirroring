package cmd

import (
	"testing"

	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/nats-io/nats.go"
)

const (
	// message to publish
	message = "snapp.cab"
)

func Test(t *testing.T) {
	// loading configs
	cfg := config.Load()

	{
		// Connect to NATS server 1 for publishing
		nc, err := nats.Connect(cfg.Nats.Nat2)
		if err != nil {
			t.Fatal(err)
		}

		// creating a jet-stream connection
		js, err := nc.JetStream()
		if err != nil {
			t.Fatal(err)
		}

		// subscribing on subject
		_, _ = js.Subscribe(cfg.Stream.SubjectName, func(msg *nats.Msg) {
			t.Logf("[NATS2][RECIEVE] %s", msg.Data)

			// make acknowledgement
			_ = msg.Ack()
		})
	}
	{
		// Connect to NATS server 1 for publishing
		nc, err := nats.Connect(cfg.Nats.Nat1)
		if err != nil {
			t.Fatal(err)
		}

		// creating a jet-stream connection
		js, err := nc.JetStream()
		if err != nil {
			t.Fatal(err)
		}

		// subscribing on subject
		_, _ = js.Subscribe(cfg.Stream.SubjectName, func(msg *nats.Msg) {
			t.Logf("[NATS1][RECIEVE] %s", msg.Data)

			// make acknowledgement
			_ = msg.Ack()
		})

		// message sending
		for i := 1; i <= cfg.Tests; i++ {
			_, err = js.Publish(cfg.Stream.SubjectName, []byte(message))
			if err != nil {
				t.Logf("[NATS1][Test %d/%d] Error: %s\n", i, cfg.Tests, err.Error())
			} else {
				t.Logf("[NATS1][Test %d/%d] OK\n", i, cfg.Tests)
			}
		}
	}
}
