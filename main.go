package main

import (
	"errors"
	"flag"

	"github.com/amirhnajafiz/j-mirror/internal/config"
	"github.com/amirhnajafiz/j-mirror/internal/handlers"
)

const (
	SVCProvider = "provider"
	SVCConsumer = "consumer"
	SVCBoot     = "boot"
)

func main() {
	var (
		ServiceTypeFlag = flag.String("service", "bootstrap", "choose from boot/consumer/provider")
		ConfigPathFlag  = flag.String("config", "config.yaml", "config file path")
		NATSHost        = flag.String("nats", "localhost:4222", "nats host for the agent")
	)

	flag.Parse()

	// load configs
	cfg := config.Load(*ConfigPathFlag)

	// create handler
	h := handlers.Handler{}

	// start bootstrap
	switch *ServiceTypeFlag {
	case SVCConsumer:
		h.Consumer(*NATSHost)
	case SVCProvider:
		h.Provider(*NATSHost)
	case SVCBoot:
		h.Bootstrap(cfg.Nats...)
	default:
		panic(errors.New("input service type is not in (provider, consumer, or boot)"))
	}
}
