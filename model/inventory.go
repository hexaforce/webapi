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

type Inventory struct {
	InventoryID int       `gorm:"column:inventory_id;primary_key" json:"inventory_id"`
	FilmID      int       `gorm:"column:film_id" json:"film_id"`
	StoreID     int       `gorm:"column:store_id" json:"store_id"`
	LastUpdate  time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (i *Inventory) TableName() string {
	return "inventory"
}
