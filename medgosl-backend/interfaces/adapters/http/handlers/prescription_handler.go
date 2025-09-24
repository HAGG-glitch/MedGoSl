package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type PrescriptionHandler struct {
	svc *services.PrescriptionService
}

func NewPrescriptionHandler(s *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{svc: s}
}

func (h *PrescriptionHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pr, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pr)
}
