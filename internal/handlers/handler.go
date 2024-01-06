package handlers

import (
	"time"

	"github.com/amirhnajafiz/jetstream-mirroring/internal/model"
)

// Handler manages testing components
type Handler struct {
	Stream           model.Stream
	ProviderInterval time.Duration
}
