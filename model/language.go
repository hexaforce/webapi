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

type Language struct {
	LanguageID int       `gorm:"column:language_id;primary_key" json:"language_id"`
	Name       string    `gorm:"column:name" json:"name"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (l *Language) TableName() string {
	return "language"
}
