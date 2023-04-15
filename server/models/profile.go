package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model

	UserID   uint
	Avatar   sql.NullString
	Birthday time.Time
}
