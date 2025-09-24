package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type patientRepo struct {
	db *gorm.DB
}

func NewPatientRepo(db *gorm.DB) domain.PatientRepository {
	return &patientRepo{db: db}
}

func (r *patientRepo) Create(ctx context.Context, p *models.Patient) error {
	if err := r.db.WithContext(ctx).Create(p).Error; err != nil {
		return fmt.Errorf("failed to create patient: %w", err)
	}
	return nil
}

func (r *patientRepo) GetByID(ctx context.Context, id uint) (*models.Patient, error) {
	var p models.Patient
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient by id %d: %w", id, err)
	}
	return &p, nil
}

func (r *patientRepo) GetByEmail(ctx context.Context, email string) (*models.Patient, error) {
	var p models.Patient
	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&p).Error; err != nil {
		return nil, fmt.Errorf("failed to get patient by email %s: %w", email, err)
	}
	return &p, nil
}

func (r *patientRepo) Update(ctx context.Context, id uint, fields map[string]interface{}) error {
	if err := r.db.WithContext(ctx).
		Model(&models.Patient{}).
		Where("id = ?", id).
		Updates(fields).Error; err != nil {
		return fmt.Errorf("failed to update patient %d: %w", id, err)
	}
	return nil
}

func (r *patientRepo) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Patient{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete patient %d: %w", id, err)
	}
	return nil
}
