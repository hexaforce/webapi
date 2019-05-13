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

type Staff struct {
	StaffID    int         `gorm:"column:staff_id;primary_key" json:"staff_id"`
	FirstName  string      `gorm:"column:first_name" json:"first_name"`
	LastName   string      `gorm:"column:last_name" json:"last_name"`
	AddressID  int         `gorm:"column:address_id" json:"address_id"`
	Picture    []byte      `gorm:"column:picture" json:"picture"`
	Email      null.String `gorm:"column:email" json:"email"`
	StoreID    int         `gorm:"column:store_id" json:"store_id"`
	Active     int         `gorm:"column:active" json:"active"`
	Username   string      `gorm:"column:username" json:"username"`
	Password   null.String `gorm:"column:password" json:"password"`
	LastUpdate time.Time   `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (s *Staff) TableName() string {
	return "staff"
}
