package model

import (
	"database/sql"
	"time"
)

type Role struct {
	ID        int            `gorm:"column:id"`
	UUID      string         `gorm:"column:uuid"`
	Kode      sql.NullString `gorm:"column:kode"`
	Nama      sql.NullString `gorm:"column:nama"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (Role) TableName() string {
	return "role"
}
