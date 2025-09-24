package models

type PaymentMethod string

const (
	OrangeMoney PaymentMethod = "orange_money"
	AfriMoney   PaymentMethod = "afri_money"
	QMoney      PaymentMethod = "q_money"
	Bank        PaymentMethod = "Bank"
)

type Payment struct {
	ID     uint          `gorm:"primaryKey" json:"id"`
	Method PaymentMethod `json:"method"`
	Phone  string        `gorm:"unique;not null" json:"phone"`
	RefID  string        `gorm:"unique;not null" json:"ref_id"`

	OrderID uint  `json:"order_id"`
	Order   Order `gorm:"constraint:OnDelete:CASCADE;" json:"order"`
}
