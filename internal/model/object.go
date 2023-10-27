package model

import (
	"database/sql"
	"time"
)

type Object struct {
	ID        int            `gorm:"column:id"`
	UUID      string         `gorm:"column:uuid"`
	Bucket    sql.NullString `gorm:"column:bucket"`
	Nama      sql.NullString `gorm:"column:nama"`
	Owner     sql.NullString `gorm:"column:owner"`
	Size      sql.NullInt64  `gorm:"column:size"`
	MimeType  sql.NullString `gorm:"column:mime_type"`
	Url       sql.NullString `gorm:"column:url"`
	Path      sql.NullString `gorm:"column:path"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (Object) TableName() string {
	return "objects"
}
