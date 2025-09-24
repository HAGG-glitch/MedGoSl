package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type medicationRepo struct {
	db *gorm.DB
}

func NewMedicationRepo(db *gorm.DB) domain.MedicationRepository {
	return &medicationRepo{db: db}
}

func (r *medicationRepo) Create(ctx context.Context, med *models.Medication) error {
	if err := r.db.WithContext(ctx).Create(med).Error; err != nil {
		return fmt.Errorf("failure in creating medication: %w", err)
	}
	return nil
}

func (r *medicationRepo) GetByID(ctx context.Context, id uint) (*models.Medication, error) {
	var med models.Medication

	if err := r.db.WithContext(ctx).First(med, id).Error; err != nil {
		return nil, fmt.Errorf("failure in getting your medication by ID %d: %w", id, err)
	}
	return &med, nil
}

func (r *medicationRepo) GetByName(ctx context.Context, name string) (*models.Medication, error) {
	var med models.Medication

	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&med).Error; err != nil {
		return nil, fmt.Errorf("failure in getting your medication by name %s: %w", name, err)
	}
	return &med, nil
}


func (r *medicationRepo) Update(ctx context.Context, id uint, fields map[string]interface{}) error{
	if err := r.db.WithContext(ctx).Model(&models.Medication{}).Where("id = ?", id).Updates(fields).Error; err != nil{
		return fmt.Errorf("failed to update medication %d: %w", id, err)
	}
	return  nil
}
