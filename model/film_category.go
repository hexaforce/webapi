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

type FilmCategory struct {
	FilmID     int       `gorm:"column:film_id;primary_key" json:"film_id"`
	CategoryID int       `gorm:"column:category_id" json:"category_id"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (f *FilmCategory) TableName() string {
	return "film_category"
}
