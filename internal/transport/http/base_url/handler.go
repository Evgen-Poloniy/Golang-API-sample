package base_url

import (
	"api_sample/internal/config"
	"api_sample/internal/service"
)

// Handler
type Handler struct {
	service service.Service
	config  *config.APIInfo
}

// NewHandler creates a new Handler with the given service
func NewHandler(service service.Service, config *config.APIInfo) *Handler {
	return &Handler{
		service: service,
		config:  config,
	}
}
