package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type PatientHandler struct {
	svc *services.PatientService
}

func NewPatientHandler(s *services.PatientService) *PatientHandler {
	return &PatientHandler{svc: s}
}

func (h *PatientHandler) Create(c *gin.Context) {
	var p models.Patient
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Create(c.Request.Context(), &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *PatientHandler) GetByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	p, err := h.svc.GetByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}
