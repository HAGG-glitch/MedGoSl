package models

type Pharmacy struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnDelete:CASCADE;" json:"-"`

	ProfilePic string  `json:"profile_pic"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`

	Medications []Medication `gorm:"foreignKey:PharmacyID" json:"medications,omitempty"`

	Orders []Order `gorm:"foreignKey:PharmacyID" json:"orders,omitempty"`

	PaymentID *uint    `json:"payment_id"`
	Payment   *Payment `gorm:"foreignKey:PaymentID" json:"payment,omitempty"`
}
