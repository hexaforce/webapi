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

type Rental struct {
	RentalID    int       `gorm:"column:rental_id;primary_key" json:"rental_id"`
	RentalDate  time.Time `gorm:"column:rental_date" json:"rental_date"`
	InventoryID int       `gorm:"column:inventory_id" json:"inventory_id"`
	CustomerID  int       `gorm:"column:customer_id" json:"customer_id"`
	ReturnDate  null.Time `gorm:"column:return_date" json:"return_date"`
	StaffID     int       `gorm:"column:staff_id" json:"staff_id"`
	LastUpdate  time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (r *Rental) TableName() string {
	return "rental"
}
