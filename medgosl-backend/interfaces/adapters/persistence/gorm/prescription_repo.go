package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type prescriptionRepo struct {
	db *gorm.DB
}


func NewPrescriptionRepo(db *gorm.DB) domain.PrescriptionRepository{
	return &prescriptionRepo{db: db}
}

func (r *prescriptionRepo)Create(ctx context.Context, pr *models.Prescription)error{
	if err := r.db.WithContext(ctx).Create(pr).Error; err != nil{
		return fmt.Errorf("failed in creating your prescirption: %w",err)
	}
	return nil
}

func (r *prescriptionRepo)GetByID(ctx context.Context, id uint)(*models.Prescription, error){
	var pr models.Prescription

	if err := r.db.WithContext(ctx).First(&pr, id).Error; err != nil{
		return nil, fmt.Errorf("failure in getting your prescription by ID %d: %w", id, err)
	}

	return &pr, nil
}

func (r *prescriptionRepo)Update(ctx context.Context,   id uint,fields map[string]interface{}) error{
	if err := r.db.WithContext(ctx).Model(&models.Prescription{}).Where("id = ?", id).Updates(fields).Error; err != nil{
		return fmt.Errorf("failure in updating your prescription %d: %w", id, err)
	}
	return  nil
}