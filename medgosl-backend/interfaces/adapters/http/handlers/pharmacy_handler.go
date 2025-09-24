package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type PharmacyHandler struct {
	svc *services.PharmacyService
}

func NewPharmacyHandler(s *services.PharmacyService) *PharmacyHandler {
	return &PharmacyHandler{svc: s}
}

func (h *PharmacyHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	ph, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ph)
}
