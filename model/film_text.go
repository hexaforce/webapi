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

type FilmText struct {
	FilmID      int         `gorm:"column:film_id;primary_key" json:"film_id"`
	Title       string      `gorm:"column:title" json:"title"`
	Description null.String `gorm:"column:description" json:"description"`
}

// TableName sets the insert table name for this struct type
func (f *FilmText) TableName() string {
	return "film_text"
}
