package dto

type CreateOrderDTO struct {
    PatientID      uint    `json:"patient_id" binding:"required"`
    PharmacyID     *uint   `json:"pharmacy_id,omitempty"`
    PrescriptionID *uint   `json:"prescription_id,omitempty"`
    Lat            float64 `json:"lat"`
    Lng            float64 `json:"lng"`
}