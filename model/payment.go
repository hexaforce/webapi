package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Payment struct {
	PaymentID   int       `gorm:"column:payment_id;primary_key" json:"payment_id"`
	CustomerID  int       `gorm:"column:customer_id" json:"customer_id"`
	StaffID     int       `gorm:"column:staff_id" json:"staff_id"`
	RentalID    null.Int  `gorm:"column:rental_id" json:"rental_id"`
	Amount      float64   `gorm:"column:amount" json:"amount"`
	PaymentDate time.Time `gorm:"column:payment_date" json:"payment_date"`
	LastUpdate  time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (p *Payment) TableName() string {
	return "payment"
}
