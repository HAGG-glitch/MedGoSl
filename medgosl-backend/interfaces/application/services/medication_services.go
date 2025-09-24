package services

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type MedicationService struct {
	medications domain.MedicationRepository
}

func NewMedicationService(m domain.MedicationRepository) *MedicationService {
	return &MedicationService{medications: m}
}

func (s *MedicationService) GetByID(ctx context.Context, id uint) (*models.Medication, error) {
	return s.medications.GetByID(ctx, id)
}
