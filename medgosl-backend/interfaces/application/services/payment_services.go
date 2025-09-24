package services

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type PaymentService struct {
	payments domain.PaymentRepository
}

func NewPaymentService(p domain.PaymentRepository) *PaymentService {
	return &PaymentService{payments: p}
}

func (s *PaymentService) GetByRef(ctx context.Context, ref string) (*models.Payment, error) {
	return s.payments.GetByRef(ctx, ref)
}
