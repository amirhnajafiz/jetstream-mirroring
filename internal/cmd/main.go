package cmd

import (
	"log"

	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/nats-io/nats.go"
)

const (
	// message to publish
	message = "snapp.cab"
)

// Execute connect to both nats servers and publish on them
func Execute(cfg config.Config) {
	{
		// Connect to NATS server 1
		nc, err := nats.Connect(cfg.Nat1)
		if err != nil {
			panic(err)
		}

		js, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}

		// create a jet-stream instance
		err = createStream(js, cfg)
		if err != nil {
			panic(err)
		}

		for {
			_, err = js.Publish(cfg.SubjectName, []byte(message))
			if err == nil {
				log.Println("Done Nats1")

				break
			}
		}
	}
	{
		// Connect to NATS server 2
		nc, err := nats.Connect(cfg.Nat2)
		if err != nil {
			panic(err)
		}

		js, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}

		// create a jet-stream instance
		err = createStream(js, cfg)
		if err != nil {
			panic(err)
		}

		for {
			_, err = js.Publish(cfg.SubjectName, []byte(message))
			if err == nil {
				log.Println("Done Nats1")

				break
			}
		}
	}
}

func createStream(js nats.JetStreamContext, cfg config.Config) error {
	stream, err := js.StreamInfo(cfg.StreamName)
	if err != nil {
		log.Println(err)
	}

	if stream == nil {
		log.Printf("creating stream %q and subjects %q", cfg.StreamName, cfg.Subject)

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
