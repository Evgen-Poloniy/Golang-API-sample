package base_url

import (
	"api_sample/internal/middleware"
	"api_sample/pkg/logger"

	// _ "api_sample/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *Handler, logger logger.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.HandleMethodNotAllowed = true

	router.Use(
		middleware.Logger(logger),
		middleware.SecureHeaders(),
		middleware.CORS(),
		gin.Recovery(),
		middleware.ErrorHandler(),
	)

	router.GET("/", handler.info)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", handler.health)

	return router
}
