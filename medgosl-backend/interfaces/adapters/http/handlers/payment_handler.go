package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type PaymentHandler struct {
	svc *services.PaymentService
}

func NewPaymentHandler(s *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{svc: s}
}

func (h *PaymentHandler) GetByRef(c *gin.Context) {
	ref := c.Param("ref")
	p, err := h.svc.GetByRef(c.Request.Context(), ref)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}
