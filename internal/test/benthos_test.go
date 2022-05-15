package test

import (
	"log"
	"testing"
	"time"

	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/nats-io/nats.go"
)

const (
	// message to publish
	message = "snapp.cab"
)

func TestBenthos(t *testing.T) {
	cfg := config.Load()

	// Connect to NATS server 1
	nc, err := nats.Connect(cfg.Nat1)
	if err != nil {
		panic(err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTimer(1 * time.Second)
	for i := range ticker.C {
		_, err = js.Publish(cfg.SubjectName, []byte(message))
		if err != nil {
			t.Errorf("[Test %d] Error: %s\n", i.Second(), err.Error())
		} else {
			log.Printf("[Test %d] Done\n", i.Second())
		}
	}
}
