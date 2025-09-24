package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type driverRepo struct {
	db *gorm.DB
}

func NewDriverRepo(db *gorm.DB) domain.DriverRepository {
	return &driverRepo{db: db}
}

func (r *driverRepo) Create(ctx context.Context, d *models.Driver) error {
	if err := r.db.WithContext(ctx).Create(d).Error; err != nil {
		return  fmt.Errorf("failed to create driver: %w", err)
	}
	return nil
}

func (r *driverRepo) GetByID(ctx context.Context, id uint) (*models.Driver, error) {
	var d models.Driver

	if err := r.db.WithContext(ctx).First(&d, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get driver by id %d: %w", id, err)
	}

	return &d, nil
}

func (r *driverRepo) GetByEmail(ctx context.Context, email string) (*models.Driver, error) {
	var d models.Driver

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&d).Error; err != nil {
		return nil, fmt.Errorf("failed to get driver by email %s: %w", email, err)
	}

	return &d, nil
}

func (r *driverRepo) FindAvailableWithin(ctx context.Context, lat, lng float64, radiusMeters int) ([]*models.Driver, error) {
	// Haversine formula in raw SQL (approx). 6371000 = earth radius in meters
	sql := `
    SELECT d.id, u.name, u.phone, d.lat, d.lng, d.available, d.updated_at
    FROM drivers d
	JOIN users u ON d.user_id = u.id
    WHERE available = true AND
    (6371000 * acos(
        cos(radians(?)) * cos(radians(lat)) * cos(radians(lng) - radians(?)) +
        sin(radians(?)) * sin(radians(lat))
    )) < ?
    ORDER BY (6371000 * acos(
        cos(radians(?)) * cos(radians(lat)) * cos(radians(lng) - radians(?)) +
        sin(radians(?)) * sin(radians(lat))
    ))
    LIMIT 10
    `
	rows, err := r.db.WithContext(ctx).Raw(sql, lat, lng, lat, radiusMeters, lat, lng, lat).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var drivers []*models.Driver
	for rows.Next() {
		var d models.Driver
		if err := rows.Scan(&d.ID, &d.User.Name, &d.User.Phone, &d.Lat, &d.Lng, &d.Available, &d.UpdatedAt); err != nil {
			return nil, err
		}
		drivers = append(drivers, &d)
	}
	return drivers, nil
}

func (r *driverRepo) UpdateLocation(ctx context.Context, id uint, lat, lng float64) error {
	if err := r.db.WithContext(ctx).Model(&models.Driver{}).Where("id = ?", id).
		Updates(map[string]interface{}{"lat": lat, "lng": lng, "updated_at": gorm.Expr("NOW()")}).Error; err != nil {
		return fmt.Errorf("failed to update driver location %d: %w", id, err)
	}

	return nil
}

