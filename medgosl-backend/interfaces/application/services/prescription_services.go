package services

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type PrescriptionService struct {
	prescriptions domain.PrescriptionRepository
}

func NewPrescriptionService(p domain.PrescriptionRepository) *PrescriptionService {
	return &PrescriptionService{prescriptions: p}
}

func (s *PrescriptionService) GetByID(ctx context.Context, id uint) (*models.Prescription, error) {
	return s.prescriptions.GetByID(ctx, id)
}
