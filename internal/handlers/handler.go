package handlers

import (
	"time"

	"github.com/amirhnajafiz/j-mirror/internal/model"
)

// Handler manages testing components
type Handler struct {
	Stream           model.Stream
	ProviderInterval time.Duration
}
