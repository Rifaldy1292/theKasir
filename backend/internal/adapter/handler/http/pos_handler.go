package http

import (
	"net/http"
	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"

	"github.com/gin-gonic/gin"
)

type PosHandler struct {
	posService port.PosService
}

func NewPosHandler(service port.PosService) *PosHandler {
	return &PosHandler{posService: service}
}

// GetProducts godoc
func (h *PosHandler) GetProducts(c *gin.Context) {
	workspaceID := c.Param("workspaceId")
	if workspaceID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "workspaceId is required"})
		return
	}

	products, err := h.posService.GetProducts(c.Request.Context(), workspaceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Always return an array (even if empty) for easier frontend consumption
	if products == nil {
		products = []domain.Product{}
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct godoc
func (h *PosHandler) CreateProduct(c *gin.Context) {
	workspaceID := c.Param("workspaceId")
	
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.WorkspaceID = workspaceID

	created, err := h.posService.CreateProduct(c.Request.Context(), &p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": created})
}

// Checkout godoc
func (h *PosHandler) Checkout(c *gin.Context) {
	workspaceID := c.Param("workspaceId")
	
	var req domain.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.WorkspaceID = workspaceID
	// In real app, extract CashierID from JWT middleware. Here we hardcode or require it in JSON.
	if req.CashierID == "" {
		req.CashierID = "usr_mock_123" // For MVP bypassing auth middleware complexity
	}

	tx, err := h.posService.Checkout(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": tx})
}
