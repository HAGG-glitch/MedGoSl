package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) domain.PaymentRepository {
	return &paymentRepo{db:db}
}


func (r *paymentRepo)Create(ctx context.Context, pay *models.Payment) error {
	if err := r.db.WithContext(ctx).Create(pay).Error; err != nil{
		return fmt.Errorf("falure in processing your payment: %w", err)
	}
	return nil
}



func (r *paymentRepo)GetByRef(ctx context.Context, ref string)(*models.Payment, error){
	var pay models.Payment
	if err := r.db.WithContext(ctx).Where("ref =", ref).First(&pay).Error; err!=nil {
		return nil, fmt.Errorf("failed to get payment by ref code %s: %w", ref, err)
	}
	return &pay, nil
}

