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

type Actor struct {
	ActorID    int       `gorm:"column:actor_id;primary_key" json:"actor_id"`
	FirstName  string    `gorm:"column:first_name" json:"first_name"`
	LastName   string    `gorm:"column:last_name" json:"last_name"`
	LastUpdate time.Time `gorm:"column:last_update" json:"last_update"`
}

// TableName sets the insert table name for this struct type
func (a *Actor) TableName() string {
	return "actor"
}
