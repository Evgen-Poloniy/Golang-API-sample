package base_url

import (
	"api_sample/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) info(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Info{
		ServiceName: h.config.ServiceName,
		Version:     h.config.Version,
		Description: h.config.Description,
		ApiDocsPath: h.config.APIDocsPath,
		Status:      "running",
	})
}
