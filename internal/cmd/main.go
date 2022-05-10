package cmd

import (
	"log"

	"github.com/nats-io/nats.go"
)

const (
	// nats servers addresses
	nats1 = "nats://0.0.0.0:4222"
	nats2 = "nats://0.0.0.0:4223"

	// stream configs
	streamName     = "snapp"
	streamSubjects = "snapp.*"
	subjectName    = "snapp.created"

	// message to publish
	message = "snapp.cab"
)

// Execute connect to both nats servers and publish on them
func Execute() {
	{
		// Connect to NATS server 1
		nc, _ := nats.Connect(nats1)
		js, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}

		// create a jet-stream instance
		err = createStream(js)
		if err != nil {
			panic(err)
		}

		for {
			_, err = js.Publish(subjectName, []byte(message))
			if err == nil {
				break
			}
		}
	}
	{
		// Connect to NATS server 2
		nc, _ := nats.Connect(nats2)
		js, err := nc.JetStream()
		if err != nil {
			log.Fatal(err)
		}

		// create a jet-stream instance
		err = createStream(js)
		if err != nil {
			panic(err)
		}

		for {
			_, err = js.Publish(subjectName, []byte(message))
			if err == nil {
				break
			}
		}
	}
}

func createStream(js nats.JetStreamContext) error {
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}

	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)

		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})

		if err != nil {
			return err
		}
	}

	return nil
}
