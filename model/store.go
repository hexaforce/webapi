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

type Store struct {
	StoreID        int       `gorm:"column:store_id;primary_key" json:"store_id"`
	ManagerStaffID int       `gorm:"column:manager_staff_id" json:"manager_staff_id"`
	AddressID      int       `gorm:"column:address_id" json:"address_id"`
	LastUpdate     time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (s *Store) TableName() string {
	return "store"
}
