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

type City struct {
	CityID     int       `gorm:"column:city_id;primary_key" json:"city_id"`
	City       string    `gorm:"column:city" json:"city"`
	CountryID  int       `gorm:"column:country_id" json:"country_id"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (c *City) TableName() string {
	return "city"
}
