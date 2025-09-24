package services

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
)

type DriverService struct {
	drivers domain.DriverRepository
}

func NewDriverService(d domain.DriverRepository) *DriverService {
	return &DriverService{drivers: d}
}

func (s *DriverService) UpdateLocation(ctx context.Context, id uint, lat, lng float64) error {
	return s.drivers.UpdateLocation(ctx, id, lat, lng)
}
