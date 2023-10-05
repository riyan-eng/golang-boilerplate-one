package model

import (
	"database/sql"
	"time"
)

type UserData struct {
	ID           int            `gorm:"column:id"`
	UUID         string         `gorm:"column:uuid"`
	Nama         sql.NullString `gorm:"column:nama"`
	NIK          sql.NullString `gorm:"column:nik"`
	NomorTelepon sql.NullString `gorm:"column:nomor_telepon"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
}

func (UserData) TableName() string {
	return "user_datas"
}
