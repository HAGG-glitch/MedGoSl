package models

type Patient struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	UserID   uint `json:"user_id"`
	User     User `gorm:"constraint:OnDelete:CASCADE;" json:"-"`

	ProfilePic string  `json:"profile_pic"`
	Age        int     `json:"age"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`

	Prescriptions []Prescription `gorm:"foreignKey:PatientID" json:"prescriptions,omitempty"`
}
