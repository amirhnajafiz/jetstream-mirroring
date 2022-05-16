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

	// Connect to NATS server
	nc, err := nats.Connect(cfg.Nats.Nat1)
	if err != nil {
		t.Fatal(err)
	}

	// creating a jet-stream connection
	js, err := nc.JetStream()
	if err != nil {
		t.Fatal(err)
	}

	// message sending
	for i := 1; i <= cfg.Tests; i++ {
		_, err = js.Publish(cfg.Stream.SubjectName, []byte(message))
		if err != nil {
			t.Logf("[Test %d/%d] Error: %s\n", i, cfg.Tests, err.Error())
		} else {
			t.Logf("[Test %d/%d] OK\n", i, cfg.Tests)
		}
	}
}
