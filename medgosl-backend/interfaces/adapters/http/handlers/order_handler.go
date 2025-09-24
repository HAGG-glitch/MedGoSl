package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/http/tracker"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/dto"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

type OrderHandler struct {
	svc *services.OrderService
	hub *tracker.TrackerHub
}

func NewOrderHandler(svc *services.OrderService, hub *tracker.TrackerHub) *OrderHandler {
	return &OrderHandler{svc: svc, hub: hub}
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req dto.CreateOrderDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	o, err := h.svc.CreateOrder(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, o)
}

func (h *OrderHandler) TrackSSE(c *gin.Context) {
	idStr := c.Param("id")
	id64, _ := strconv.ParseUint(idStr, 10, 32)
	id := uint(id64)
	ch := h.hub.Subscribe(id)
	defer h.hub.Unsubscribe(id, ch)

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.WriteHeader(http.StatusOK)

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()

	// Optionally send a heartbeat
	heartbeat := time.NewTicker(20 * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case upd, ok := <-ch:
			if !ok {
				return
			}
			b, _ := json.Marshal(upd)
			fmt.Fprintf(c.Writer, "data: %s\n\n", b)
			flusher.Flush()
		case <-heartbeat.C:
			fmt.Fprintf(c.Writer, ":\n\n") // comment line to keep connection open
			flusher.Flush()
		}
	}
}

func (h *OrderHandler) PostDriverLocation(c *gin.Context) {
	// body: { "driver_id": 1, "order_id": 2, "lat": -0.123, "lng": 11.22 }
	type body struct {
		DriverID uint    `json:"driver_id"`
		OrderID  uint    `json:"order_id"`
		Lat      float64 `json:"lat"`
		Lng      float64 `json:"lng"`
	}
	var b body
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update driver location in DB via driver repo (omitted here; call service)
	// publish to hub
	h.hub.Publish(b.OrderID, dto.LocationUpdate{
		OrderID: b.OrderID, DriverID: b.DriverID, Lat: b.Lat, Lng: b.Lng, At: time.Now(),
	})
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *OrderHandler) ConfirmByPharmacy(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}
	var body struct {
		PharmacyID uint `json:"pharmacy_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.ConfirmOrderByPharmacy(ctx, uint(orderID), body.PharmacyID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "confirmed"})
}

func (h *OrderHandler) AssignDrive(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	if err := h.svc.AssignDriver(ctx, uint(orderID)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Assigned"})

}

func (h *OrderHandler) PickUp(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var body struct {
		DriverID uint `json:"driver_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.MarkPickedUp(ctx, uint(orderID), body.DriverID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Picked up"})

}

func (h *OrderHandler) Paid(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var body struct {
		PaymentID uint `json:"payment_id" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.MarkPaid(ctx, uint(orderID), body.PaymentID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Paid"})

}

func (h *OrderHandler) Delivered(ctx *gin.Context) {
	orderID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var body struct {
		DriverID uint   `json:"payment_id" binding:"required"`
		Ticket   string `json:"ticket" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.MarkDelivered(ctx, uint(orderID),body.DriverID, body.Ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{"status": "Delivered"})
}
