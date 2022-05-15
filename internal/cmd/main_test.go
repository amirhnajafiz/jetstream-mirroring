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
	cfg := config.Load()

	// Connect to NATS server
	nc, err := nats.Connect(cfg.Nat1)
	if err != nil {
		t.Fatal(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i < 11; i++ {
		_, err = js.Publish(cfg.SubjectName, []byte(message))
		if err != nil {
			t.Logf("[Test %d] Error: %s\n", i, err.Error())
		} else {
			t.Logf("[Test %d] Done\n", i)
		}
	}
}
