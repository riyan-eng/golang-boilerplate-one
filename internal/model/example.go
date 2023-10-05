package model

import (
	"database/sql"
	"time"
)

type Example struct {
	ID        int            `gorm:"column:id"`
	UUID      string         `gorm:"column:uuid"`
	Nama      string         `gorm:"column:nama"`
	Detail    sql.NullString `gorm:"column:detail"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (Example) TableName() string {
	return "example"
}
