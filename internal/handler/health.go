package handler

import (
	"net/http"
	"project/internal/dto"

	"github.com/gin-gonic/gin"
)

// health godoc
// @Summary Health check
// @Description Check service availability
// @Tags system
// @Success 200 {object} dto.Response
// @Router /health [get]
func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Health{Status: "ok"})
}
