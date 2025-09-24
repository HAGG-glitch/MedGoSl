package models

import "time"

type Medication struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	MedicationImage string    `json:"medication_image"`
	Description     string    `json:"description"`
	Price           float64   `json:"price"`
	Stock           int       `json:"stock"`
	DateMade        time.Time `json:"date_made"`
	ExpiryDate      time.Time `json:"expiry_date"`

	PharmacyID uint     `json:"pharmacy_id"`
	Pharmacy   Pharmacy `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
