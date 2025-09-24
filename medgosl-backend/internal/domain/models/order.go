package models

import "time"

type OrderStatus string

const (
	StatusPending    OrderStatus = "pending"    // created but not confirmed
	StatusConfirmed  OrderStatus = "confirmed"  // pharmacy verified
	StatusAssigned   OrderStatus = "assigned"   // driver assigned
	StatusPickedUp   OrderStatus = "picked_up"  // driver collected from pharmacy
	StatusPaymentConfirmed OrderStatus = "payment_confirmed"// confirms paymentm
	StatusDelivered  OrderStatus = "delivered"  // completed
	StatusCancelled  OrderStatus = "cancelled"  // cancelled
)

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Status    OrderStatus `json:"status"`
	Lat       float64     `json:"lat"`
	Lng       float64     `json:"lng"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`

	//  Relations
	PatientID   uint     `json:"patient_id"`
	Patient     *Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	PrescriptionID *uint         `json:"prescription_id"`
	Prescription   *Prescription `gorm:"foreignKey:PrescriptionID" json:"prescription,omitempty"`

	PharmacyID *uint     `json:"pharmacy_id"`
	Pharmacy   *Pharmacy `gorm:"foreignKey:PharmacyID" json:"pharmacy,omitempty"`

	DriverID *uint   `json:"driver_id"`
	Driver   *Driver `gorm:"foreignKey:DriverID" json:"driver,omitempty"`

	PaymentID *uint    `json:"payment_id"`
	Payment   *Payment `gorm:"foreignKey:PaymentID" json:"payment,omitempty"`

	ConfirmationTicket string `json:"confirmation_ticket,omitempty"`
}
