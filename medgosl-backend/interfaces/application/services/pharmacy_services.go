package services

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type PharmacyService struct {
	pharmacies domain.PharmacyRepository
}

func NewPharmacyService(p domain.PharmacyRepository) *PharmacyService {
	return &PharmacyService{pharmacies: p}
}

func (s *PharmacyService) GetByID(ctx context.Context, id uint) (*models.Pharmacy, error) {
	return s.pharmacies.GetByID(ctx, id)
}
