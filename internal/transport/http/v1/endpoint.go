package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getDataByUnitGUID godoc
// @Summary Get parsed data by Unit GUID
// @Description Returns paginated parsed data for a given unit GUID
// @Tags data
// @Produce json
// @Param unit_guid query string true "Unit GUID"
// @Param limit query int false "Limit of records (1-1000)" default(1000) minimum(1) maximum(1000)
// @Param page query int false "Page number (>=1)" default(1) minimum(1)
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Error "Invalid request parameters"
// @Failure 500 {object} dto.Error "Internal server error"
// @Security ApiKeyAuth
// @Router /get-data [get]
func (h *Handler) endpoint(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
