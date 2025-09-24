package main

import (
	"log"

	"fmt"

	"github.com/HAGG-glitch/MedGoSl.git/configs"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/http"

	"github.com/HAGG-glitch/MedGoSl.git/internal/domain/models"

	"github.com/joho/godotenv"
	

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found, using system environment")
	}
}


func main() {
	

	// Load config
	cfg := configs.Load()

	// Connect to DB
	db, err := gorm.Open(postgres.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Driver{},
		&models.Pharmacy{},
		&models.Medication{},
		&models.Prescription{},
		&models.Payment{},
		&models.Order{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println(" Database connected and migrated successfully")

	// Setup router
	r := http.SetupRouter(db, cfg)

	// Run server
	if err := r.Run(cfg.Addr); err != nil {
		log.Fatal(" Failed to start server:", err)
	}
}
