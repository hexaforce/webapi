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

type Country struct {
	CountryID  int       `gorm:"column:country_id;primary_key" json:"country_id"`
	Country    string    `gorm:"column:country" json:"country"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (c *Country) TableName() string {
	return "country"
}
