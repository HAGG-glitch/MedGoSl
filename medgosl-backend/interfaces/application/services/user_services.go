package services

import (
	"context"
	"errors"
	// "errors"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain"
	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	users domain.UserRepository
}

func NewUserService(u domain.UserRepository) *UserService {
	return &UserService{users: u}
}

func (s *UserService) Register(ctx context.Context, u *models.User) error {
	return s.users.Create(ctx, u)
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*models.User, error) {
	return s.users.GetByID(ctx, id)
}

func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.users.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return u, nil
}

func (s *UserService) Update(ctx context.Context, id uint, fields map[string]interface{}) error {
	return s.users.Update(ctx, id, fields)
}
