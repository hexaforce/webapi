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

type Address struct {
	AddressID  int         `gorm:"column:address_id;primary_key" json:"address_id"`
	Address    string      `gorm:"column:address" json:"address"`
	Address2   null.String `gorm:"column:address2" json:"address2"`
	District   string      `gorm:"column:district" json:"district"`
	CityID     int         `gorm:"column:city_id" json:"city_id"`
	PostalCode null.String `gorm:"column:postal_code" json:"postal_code"`
	Phone      string      `gorm:"column:phone" json:"phone"`
	LastUpdate time.Time   `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (a *Address) TableName() string {
	return "address"
}
