package domain

import (
	"context"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
)

type OrderRepository interface {
	Create(ctx context.Context, o *models.Order) error
	GetByID(ctx context.Context, id uint) (*models.Order, error)
	Save(ctx context.Context, o *models.Order) error
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
}

type DriverRepository interface {
	Create(ctx context.Context, d *models.Driver) error
	GetByID(ctx context.Context, id uint) (*models.Driver, error)
	GetByEmail(ctx context.Context, email string) (*models.Driver, error)
	FindAvailableWithin(ctx context.Context, lat, lng float64, radiusMeters int) ([]*models.Driver, error)
	UpdateLocation(ctx context.Context, id uint, lat, lng float64) error
}

type PatientRepository interface {
	Create(ctx context.Context, p *models.Patient) error
	GetByID(ctx context.Context, id uint) (*models.Patient, error)
	GetByEmail(ctx context.Context, email string) (*models.Patient, error)
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
	Delete(ctx context.Context, id uint) error
}

type PharmacyRepository interface {
	Create(ctx context.Context, ph *models.Pharmacy) error
	GetByID(ctx context.Context, id uint) (*models.Pharmacy, error)
	GetByEmail(ctx context.Context, email string) (*models.Pharmacy, error)
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
}

type MedicationRepository interface{
	Create(ctx context.Context, med *models.Medication)error
	GetByID(ctx context.Context, id uint) (*models.Medication, error)
	GetByName(ctx context.Context, name string) (*models.Medication, error)
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
}

type PrescriptionRepository interface {
	Create(ctx context.Context, pr *models.Prescription) error
	GetByID(ctx context.Context, id uint) (*models.Prescription, error)
	Update(ctx context.Context,  id uint, fields map[string]interface{}) error
}

type PaymentRepository interface {
	Create(ctx context.Context, pay *models.Payment) error
	GetByRef(ctx context.Context, ref string) (*models.Payment, error)
}

type UserRepository interface {
	Create(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
}
