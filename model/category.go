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

type Category struct {
	CategoryID int       `gorm:"column:category_id;primary_key" json:"category_id"`
	Name       string    `gorm:"column:name" json:"name"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "category"
}
