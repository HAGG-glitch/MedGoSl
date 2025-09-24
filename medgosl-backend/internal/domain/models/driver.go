package models

import "time"

type Driver struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	UserID   uint `json:"user_id"`
	User     User `gorm:"constraint:OnDelete:CASCADE;" json:"-"`

	ProfilePic string    `json:"profile_pic"`
	Lat        float64   `json:"lat"`
	Lng        float64   `json:"lng"`
	Available  bool      `json:"available"`
	UpdatedAt  time.Time `json:"updated_at"`

	Orders []Order `gorm:"foreignKey:DriverID" json:"orders,omitempty"`
}	
