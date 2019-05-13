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

type Film struct {
	FilmID             int         `gorm:"column:film_id;primary_key" json:"film_id"`
	Title              string      `gorm:"column:title" json:"title"`
	Description        null.String `gorm:"column:description" json:"description"`
	LanguageID         int         `gorm:"column:language_id" json:"language_id"`
	OriginalLanguageID null.Int    `gorm:"column:original_language_id" json:"original_language_id"`
	RentalDuration     int         `gorm:"column:rental_duration" json:"rental_duration"`
	RentalRate         float64     `gorm:"column:rental_rate" json:"rental_rate"`
	Length             null.Int    `gorm:"column:length" json:"length"`
	ReplacementCost    float64     `gorm:"column:replacement_cost" json:"replacement_cost"`
	Rating             null.String `gorm:"column:rating" json:"rating"`
	SpecialFeatures    null.String `gorm:"column:special_features" json:"special_features"`
	LastUpdate         time.Time   `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (f *Film) TableName() string {
	return "film"
}
