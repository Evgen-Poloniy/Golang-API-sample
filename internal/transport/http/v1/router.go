package v1

import (
	"api_sample/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, handler *Handler, apiKeyHash []byte) {
	v1 := router.Group("/api/v1")
	v1.Use(
		middleware.APIKeyAuth(apiKeyHash),
	)
	{
		v1.GET("/", handler.endpoint)
	}
}
