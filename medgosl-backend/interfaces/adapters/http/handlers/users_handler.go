package handlers


import (
	"context"
	"net/http"
	"time"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/dto"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(u *services.UserService) *UserHandler {
	return &UserHandler{userService: u}
}

// POST /users/register
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password, // ⚠️ Hash before saving in real app
		Role:     models.UserType(req.Role),
		Phone:    req.Phone,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.userService.Register(ctx, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

// POST /users/login
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := h.userService.Login(ctx, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// 🔑 TODO: Generate JWT here
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"user":    user,
	})
}
