package gormrepo

import (
	"context"
	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, u *models.User) error {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return fmt.Errorf("failed in creating user account: %w", err)
	}
	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var u models.User

	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, fmt.Errorf("failure in getting user by ID %d: %w", id,err)
	}

	return &u, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var u models.User

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, fmt.Errorf("failure in getting user by email %s: %w", email,err)
	}

	return &u, nil
}

func (r *userRepo) Update(ctx context.Context,  id uint, fields map[string]interface{}) error {
	if err := r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(fields).Error; err != nil{
		return fmt.Errorf("failure in updating your user details %d: %w", id, err)
	}
	return nil
}
