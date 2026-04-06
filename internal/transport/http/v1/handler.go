package v1

import (
	"api_sample/internal/service"
)

// Handler
type Handler struct {
	service service.Service
}

// NewHandler creates a new Handler with the given service
func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
