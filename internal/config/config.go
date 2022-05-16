package config

import (
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

// Config main configs of the application
type Config struct {
	Nats   Nats   `koanf:"nats"`
	Stream Stream `koanf:"stream"`
	Tests  int    `koanf:"number_of_tests"`
}

// Nats configs for our nats servers
type Nats struct {
	Nat1 string `koanf:"nat1_url"`
	Nat2 string `koanf:"nat2_url"`
}

// Stream configs for nats streams
type Stream struct {
	StreamName  string `koanf:"stream_name"`
	Subject     string `koanf:"subject"`
	SubjectName string `koanf:"subject_name"`
}

// Load loads all of our configs
func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		_ = fmt.Errorf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		_ = fmt.Errorf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %v\n", err)
	}

	return instance
}
