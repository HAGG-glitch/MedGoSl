package models

type Prescription struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Phone       string `json:"phone"`
	Notes       string `json:"notes"`
	UploadImage string `json:"upload_image"`

	PatientID uint    `json:"patient_id"`
	Patient   Patient `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}
