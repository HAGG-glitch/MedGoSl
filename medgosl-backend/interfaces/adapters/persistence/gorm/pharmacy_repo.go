package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type pharmacyRepo struct {
	db *gorm.DB
}

func NewPharmacy(db *gorm.DB) domain.PharmacyRepository {
	return &pharmacyRepo{db: db}
}

func (r *pharmacyRepo) Create(ctx context.Context, ph *models.Pharmacy) error {
	// return

	if err := r.db.WithContext(ctx).Create(ph).Error; err != nil {
		return fmt.Errorf("failed in creating pharmacy account: %w", err)
	}

	return nil
}

func (r *pharmacyRepo) GetByID(ctx context.Context, id uint) (*models.Pharmacy, error) {
	var ph models.Pharmacy

	if err := r.db.WithContext(ctx).First(&ph, id).Error; err != nil {
		return nil, fmt.Errorf("failed in getting pharmacy by id %d: %w", id, err)
	}

	return &ph, nil

}

func (r *pharmacyRepo) GetByEmail(ctx context.Context, email string) (*models.Pharmacy, error) {
	var ph models.Pharmacy

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&ph).Error; err != nil {
		return nil, fmt.Errorf("failed in getting pharmacy by email %s: %w", email, err)
	}

	return &ph, nil
}

func (r *pharmacyRepo) Update(ctx context.Context,  id uint, fields map[string]interface{}) error {

	if err := r.db.WithContext(ctx).Model(&models.Pharmacy{}).Where("id = ?", id).Updates(fields).Error; err != nil {
		return fmt.Errorf("failed in updating the pharmacy %d: %w", id, err)
	}

	return nil
}
