package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type MedicationHandler struct {
	svc *services.MedicationService
}

func NewMedicationHandler(s *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{svc: s}
}

func (h *MedicationHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	m, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, m)
}
