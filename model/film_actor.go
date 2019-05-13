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

type FilmActor struct {
	ActorID    int       `gorm:"column:actor_id;primary_key" json:"actor_id"`
	FilmID     int       `gorm:"column:film_id" json:"film_id"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (f *FilmActor) TableName() string {
	return "film_actor"
}
