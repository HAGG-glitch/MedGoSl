package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type DriverHandler struct {
	svc *services.DriverService
}

func NewDriverHandler(s *services.DriverService) *DriverHandler {
	return &DriverHandler{svc: s}
}

func (h *DriverHandler) UpdateLocation(c *gin.Context) {
	type req struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	var body req
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.svc.UpdateLocation(c.Request.Context(), uint(id), body.Lat, body.Lng); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
