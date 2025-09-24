package services

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type PatientService struct {
	patients domain.PatientRepository
}

func NewPatientService(p domain.PatientRepository) *PatientService {
	return &PatientService{patients: p}
}

func (s *PatientService) Create(ctx context.Context, p *models.Patient) error {
	return s.patients.Create(ctx, p)
}

func (s *PatientService) GetByID(ctx context.Context, id uint) (*models.Patient, error) {
	return s.patients.GetByID(ctx, id)
}
