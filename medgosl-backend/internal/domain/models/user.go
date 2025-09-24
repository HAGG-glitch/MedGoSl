package models

type UserType string

const (
	UserAdmin    UserType = "admin"
	UserCustomer UserType = "customer"
	UserDriver   UserType = "driver"
	UserPharmacy UserType = "pharmacy"
)

type User struct {
	ID             uint     `gorm:"primaryKey" json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email" gorm:"unique;not null"`
	Password       string   `json:"password,omitempty"`         // plain (for registration only)
	PasswordHashed string   `json:"password_hashed,omitempty"` // hashed password
	Role           UserType `json:"role"`
	Phone          string   `json:"phone" gorm:"unique;not null"`

	// Relations
	Patient  *Patient  `gorm:"foreignKey:UserID" json:"patient,omitempty"`
	Driver   *Driver   `gorm:"foreignKey:UserID" json:"driver,omitempty"`
	Pharmacy *Pharmacy `gorm:"foreignKey:UserID" json:"pharmacy,omitempty"`
}
