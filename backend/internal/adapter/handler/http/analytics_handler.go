package http

import (
	"net/http"
	"thekasir/internal/core/port"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	service port.AnalyticsService
}

func NewAnalyticsHandler(service port.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{service: service}
}

func (h *AnalyticsHandler) GetSummary(c *gin.Context) {
	workspaceID := c.Param("workspaceId")
	if workspaceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "workspaceId is required"})
		return
	}

	data, err := h.service.GetDashboardData(c.Request.Context(), workspaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
