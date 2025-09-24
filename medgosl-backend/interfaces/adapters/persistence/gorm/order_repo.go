package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) domain.OrderRepository {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Create(ctx context.Context, o *models.Order) error {
	if err := r.db.WithContext(ctx).Create(o).Error; err != nil {
		return fmt.Errorf("failure in creating your order: %w", err)
	}
	return nil
}

func (r *OrderRepo) GetByID(ctx context.Context, id uint) (*models.Order, error) {
	var o models.Order
	if err := r.db.WithContext(ctx).First(&o, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get order by id %d: %w", id, err)
	}
	return &o, nil
}

func (r *OrderRepo) Save(ctx context.Context, o *models.Order) error {
	if err := r.db.WithContext(ctx).Save(o).Error; err != nil {
		return fmt.Errorf("failed to save order %d: %w", o.ID, err)
	}
	return nil
}

func (r *OrderRepo) Update(ctx context.Context, id uint, fields map[string]interface{}) error {
	if err := r.db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", id).Updates(fields).Error; err != nil {
		return fmt.Errorf("failed to update patient %d: %w", id, err)
	}
	return nil

}
