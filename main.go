package main

import (
	"errors"
	"flag"
	"log"
	"time"

	"github.com/amirhnajafiz/jetstream-mirroring/internal/config"
	"github.com/amirhnajafiz/jetstream-mirroring/internal/handlers"
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
	h := handlers.Handler{
		Stream:           cfg.Stream,
		ProviderInterval: time.Duration(cfg.Interval) * time.Second,
	}

	// start bootstrap
	switch *ServiceTypeFlag {
	case SVCConsumer:
		if err := h.Consumer(*NATSHost); err != nil {
			panic(err)
		}
	case SVCProvider:
		if err := h.Provider(*NATSHost); err != nil {
			panic(err)
		}
	case SVCBoot:
		if err := h.Bootstrap(cfg.Clusters); len(err) > 0 {
			for _, msg := range err {
				log.Println(msg)
			}

			panic(errors.New("failed to bootstrap all clusters"))
		}
	default:
		panic(errors.New("input service type is not in (provider, consumer, or boot)"))
	}
}
