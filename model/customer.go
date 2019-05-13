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

type Customer struct {
	CustomerID int         `gorm:"column:customer_id;primary_key" json:"customer_id"`
	StoreID    int         `gorm:"column:store_id" json:"store_id"`
	FirstName  string      `gorm:"column:first_name" json:"first_name"`
	LastName   string      `gorm:"column:last_name" json:"last_name"`
	Email      null.String `gorm:"column:email" json:"email"`
	AddressID  int         `gorm:"column:address_id" json:"address_id"`
	Active     int         `gorm:"column:active" json:"active"`
	CreateDate time.Time   `gorm:"column:create_date" json:"create_date"`
	LastUpdate time.Time   `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (c *Customer) TableName() string {
	return "customer"
}
