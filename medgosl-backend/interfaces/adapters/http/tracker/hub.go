package tracker

import (
	"sync"
	"time"

	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/dto"
)

// TrackerHub is the concrete implementation used in handlers
type TrackerHub struct {
	mu   sync.RWMutex
	subs map[uint]map[chan dto.LocationUpdate]struct{}
}

func NewTrackerHub() *TrackerHub {
	return &TrackerHub{
		subs: make(map[uint]map[chan dto.LocationUpdate]struct{}),
	}
}

func (h *TrackerHub) Subscribe(orderID uint) chan dto.LocationUpdate {
	ch := make(chan dto.LocationUpdate, 10)
	h.mu.Lock()
	if _, ok := h.subs[orderID]; !ok {
		h.subs[orderID] = make(map[chan dto.LocationUpdate]struct{})
	}
	h.subs[orderID][ch] = struct{}{}
	h.mu.Unlock()
	return ch
}

func (h *TrackerHub) Unsubscribe(orderID uint, ch chan dto.LocationUpdate) {
	h.mu.Lock()
	if m, ok := h.subs[orderID]; ok {
		delete(m, ch)
		close(ch)
		if len(m) == 0 {
			delete(h.subs, orderID)
		}
	}
	h.mu.Unlock()
}

func (h *TrackerHub) Publish(orderID uint, upd dto.LocationUpdate) {
	h.mu.RLock()
	if m, ok := h.subs[orderID]; ok {
		for ch := range m {
			select {
			case ch <- upd:
			default: // drop if channel is blocked
			}
		}
	}
	h.mu.RUnlock()
}

func (h *TrackerHub) PublishAssignment(orderID, driverID uint) {
	h.Publish(orderID, dto.LocationUpdate{
		OrderID:  orderID,
		DriverID: driverID,
		Lat:      0,
		Lng:      0,
		At:       time.Now(),
	})
}
