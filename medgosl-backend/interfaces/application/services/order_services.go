package services

import (
	"context"
	"fmt"
	"time"

	"github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/external/googlemaps"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/dto"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type TrackerHubInterface interface {
	Publish(orderID uint, upd dto.LocationUpdate)
	PublishAssignment(orderID, driverID uint)
}

type OrderService struct {
	orders     domain.OrderRepository
	drivers    domain.DriverRepository
	pharmacies domain.PharmacyRepository
	maps       *googlemaps.Client
	hub        TrackerHubInterface
}

func NewOrderService(o domain.OrderRepository, d domain.DriverRepository, p domain.PharmacyRepository, m *googlemaps.Client, h TrackerHubInterface) *OrderService {
	return &OrderService{orders: o, drivers: d, pharmacies: p, maps: m, hub: h}
}


func (s *OrderService) CreateOrder(ctx context.Context, dto dto.CreateOrderDTO) (*models.Order, error) {
	o := &models.Order{
		PatientID:      dto.PatientID,
		PrescriptionID: dto.PrescriptionID,
		PharmacyID:     dto.PharmacyID,
		Status:         models.StatusPending,
	}
	if dto.Lat == 0 && dto.Lng == 0 {
		if dto.PharmacyID != nil {
			pharmacy, err := s.pharmacies.GetByID(ctx, *dto.PharmacyID)
			if err == nil {
				o.Lat = pharmacy.Lat
				o.Lng = pharmacy.Lng
			}
		}
	} else {
		o.Lat = dto.Lat
		o.Lng = dto.Lng
	}

	if err := s.orders.Create(ctx, o); err != nil {
		return nil, err
	}

	// find nearest drivers (within 5km)
	drivers, err := s.drivers.FindAvailableWithin(ctx, o.Lat, o.Lng, 5000)
	if err == nil && len(drivers) > 0 {
		d := drivers[0]
		o.DriverID = &d.ID
		o.Status = models.StatusPending
		if err := s.orders.Save(ctx, o); err == nil {
			s.hub.PublishAssignment(o.ID, d.ID)
		}
	}

	return o, nil
}

// confirms Order by pharmacy

func (s *OrderService) ConfirmOrderByPharmacy(ctx context.Context, orderID uint, pharmacyID uint) error {
	o, err := s.orders.GetByID(ctx, orderID)
	if err != nil {
		return err
	}
	if o.Status != models.StatusPending {
		return fmt.Errorf("order %d cannot be confirmed; current status: %s", o.ID, o.Status)
	}

	o.Status = models.StatusConfirmed
	o.PharmacyID = &pharmacyID
	o.UpdatedAt = time.Now()

	return s.orders.Update(ctx, orderID, map[string]interface{}{
		"status": models.StatusConfirmed,
		
		"updated_at": time.Now(),
	})
}

func (s *OrderService) AssignDriver(ctx context.Context, orderID uint) error {
	o, err := s.orders.GetByID(ctx, orderID)
	if err != nil {
		return err
	}
	if o.Status != models.StatusConfirmed {
		return fmt.Errorf("order %d cannot be assigned; current status: %s", o.ID, o.Status)
	}

	drivers, err := s.drivers.FindAvailableWithin(ctx, o.Lat, o.Lng, 5000)
	if err != nil || len(drivers) == 0 {
		return fmt.Errorf("no available drivers")
	}

	d := drivers[0]
	o.DriverID = &d.ID
	o.Status = models.StatusAssigned
	o.UpdatedAt = time.Now()

	if err := s.orders.Update(ctx, orderID, map[string]interface{}{
		"status": models.StatusAssigned,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}

	// publish assignment
	s.hub.PublishAssignment(o.ID, d.ID)
	return nil
}

func (s *OrderService) MarkPickedUp(ctx context.Context, orderID uint, driverID uint) error {
	o, err := s.orders.GetByID(ctx, orderID)
	if err != nil {
		return err
	}
	if o.Status != models.StatusAssigned || o.DriverID == nil || *o.DriverID != driverID {
		return fmt.Errorf("order %d cannot be marked picked up", o.ID)
	}

	o.Status = models.StatusPickedUp
	o.UpdatedAt = time.Now()

	return s.orders.Update(ctx, orderID, map[string]interface{}{
		"status": models.StatusPickedUp,
		"updated_at": time.Now(),
	})
}

func (s *OrderService) MarkPaid(ctx context.Context, orderID uint, paymentID uint) error {
	o, err := s.orders.GetByID(ctx, orderID)
	if err != nil {
		return err
	}
	if o.Status != models.StatusPickedUp {
		return fmt.Errorf("order %d cannot be paid; current status: %s", o.ID, o.Status)
	}


	o.Status = models.StatusPaymentConfirmed
	o.PaymentID = &paymentID
	o.UpdatedAt = time.Now()

	return s.orders.Update(ctx, orderID, map[string]interface{}{
		"status": models.StatusPaymentConfirmed,
		"updated_at": time.Now(),
	})
}

func (s *OrderService) MarkDelivered(ctx context.Context, orderID uint, driverID uint, ticket string) error {
	o, err := s.orders.GetByID(ctx, orderID)
	if err != nil {
		return err
	}
	if o.Status != models.StatusPaymentConfirmed || o.DriverID == nil || *o.DriverID != driverID {
		return fmt.Errorf("order %d cannot be delivered", o.ID)
	}

	o.Status = models.StatusDelivered
	o.ConfirmationTicket = ticket
	o.UpdatedAt = time.Now()

	return s.orders.Update(ctx, orderID, map[string]interface{}{
		"status": models.StatusDelivered,
		"updated_at": time.Now(),
		"confirmation_ticket": ticket,
	})
}




