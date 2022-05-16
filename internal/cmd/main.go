package cmd

import (
	"log"

	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/nats-io/nats.go"
)

// Execute connect to both nats servers and creating the streams
func Execute() {
	cfg := config.Load()

	{
		// Connect to NATS server 1
		nc, err := nats.Connect(cfg.Nats.Nat1)
		if err != nil {
			panic(err)
		}

		// creating a jet-stream connection
		js, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}

		// create a jet-stream instance
		err = createStream(js, cfg.Stream)
		if err != nil {
			panic(err)
		}

		log.Printf("[OK] first js server streams created\n")
	}
	{
		// Connect to NATS server 2
		nc, err := nats.Connect(cfg.Nats.Nat2)
		if err != nil {
			panic(err)
		}

		// creating a jet-stream connection
		js, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}

		// create a jet-stream instance
		err = createStream(js, cfg.Stream)
		if err != nil {
			panic(err)
		}

		log.Printf("[OK] second js server streams created\n")
	}
}

// this method creates our stream in js server
func createStream(js nats.JetStreamContext, cfg config.Stream) error {
	stream, err := js.StreamInfo(cfg.StreamName)
	if err != nil {
		log.Println(err)
	}

	if stream == nil {
		log.Printf("[OK] creating stream %q and subjects %q\n", cfg.StreamName, cfg.Subject)

		_, err = js.AddStream(&nats.StreamConfig{
			Name:     cfg.StreamName,
			Subjects: []string{cfg.Subject},
		})

		if err != nil {
			return err
		}
	}

	return nil
}
