package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/HAGG-glitch/MedGoSl.git/configs"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/external/googlemaps"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/http/handlers"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/http/tracker"
	gormrepo "github.com/HAGG-glitch/MedGoSl.git/interfaces/adapters/persistence/gorm"
	"github.com/HAGG-glitch/MedGoSl.git/interfaces/application/services"
)

func SetupRouter(db *gorm.DB, cfg *configs.Config) *gin.Engine {
	r := gin.Default()

	// external links
	mapsClient := googlemaps.NewClient(cfg.GoogleMapsAPIKey)

	// Init repositories
	userRepo := gormrepo.NewUserRepo(db)
	orderRepo := gormrepo.NewOrderRepo(db)
	driverRepo := gormrepo.NewDriverRepo(db)
	pharmacyRepo := gormrepo.NewPharmacy(db)
	patientRepo := gormrepo.NewPatientRepo(db)
	medicationRepo := gormrepo.NewMedicationRepo(db)
	paymentRepo := gormrepo.NewPaymentRepo(db)
	prescriptionRepo := gormrepo.NewPrescriptionRepo(db)

	// Tracker
	hub := tracker.NewTrackerHub()

	// Init services
	userService := services.NewUserService(userRepo)
	driverService := services.NewDriverService(driverRepo)
	patientService := services.NewPatientService(patientRepo)
	medicationService := services.NewMedicationService(medicationRepo)
	pharmacyService := services.NewPharmacyService(pharmacyRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	prescriptionService := services.NewPrescriptionService(prescriptionRepo)
	orderService := services.NewOrderService(orderRepo, driverRepo, pharmacyRepo, mapsClient, hub)

	// Init handlers
	userHandler := handlers.NewUserHandler(userService)
	orderHandler := handlers.NewOrderHandler(orderService, hub)
	driverHandler := handlers.NewDriverHandler(driverService)
	patientHandler := handlers.NewPatientHandler(patientService)
	medicationHandler := handlers.NewMedicationHandler(medicationService)
	pharmacyHandler := handlers.NewPharmacyHandler(pharmacyService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	prescriptionHandler := handlers.NewPrescriptionHandler(prescriptionService)

	// Routes
	api := r.Group("MedgoSL/api/v1")
	{
		api.POST("/users/register", userHandler.Register)
		api.POST("/users/login", userHandler.Login)
		// add other handlers here (orders, pharmacy, etc.)

		api.POST("/orders", orderHandler.Create)
		api.GET("/orders/:id/track", orderHandler.TrackSSE)
		api.POST("/orders/driver-location", orderHandler.PostDriverLocation)
		// (soon: confirm, assign, mark picked up, mark paid, mark delivered)
		api.PATCH("/orders/confirm_order", orderHandler.ConfirmByPharmacy)
		api.PATCH("/orders/assign_driver", orderHandler.AssignDrive)
		api.PATCH("/orders/pick_up", orderHandler.PickUp)
		api.PATCH("/orders/paid", orderHandler.Paid)
		api.PATCH("/orders/delivered", orderHandler.Delivered)
		api.POST("/patients", patientHandler.Create)
		api.GET("/patients/:id", patientHandler.GetByID)

		api.PATCH("/drivers/:id/location", driverHandler.UpdateLocation)

		api.GET("/medications/:id", medicationHandler.GetByID)
		api.GET("/pharmacies/:id", pharmacyHandler.GetByID)
		api.GET("/prescriptions/:id", prescriptionHandler.GetByID)

		api.GET("/payments/:ref", paymentHandler.GetByRef)

	}

	return r
}
