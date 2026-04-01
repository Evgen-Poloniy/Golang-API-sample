package handler

import (
	"project/internal/config"
	"project/internal/middleware"
	"project/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "project/docs"
)

// Handler
type Handler struct {
	service service.ServiceInterface
}

// NewHandler creates a new Handler with the given service
func NewHandler(service service.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

// InitHandlers initializes the HTTP handlers, routes, and middleware.
func (h *Handler) InitHandlers(config *config.LoggerConfig, apiKeyHash []byte) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// Is used logger in middleware
	router.Use(middleware.InitLoggingMiddleware(config))
	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	v1.GET("/health", h.health)
	v1.Use(middleware.APIKeyAuth(apiKeyHash))
	{

	}

	return router
}
