package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int            `gorm:"column:id"`
	UUID      string         `gorm:"column:uuid"`
	Email     sql.NullString `gorm:"column:email"`
	Password  sql.NullString `gorm:"column:password"`
	Role      sql.NullString `gorm:"column:role"`
	UserData  sql.NullString `gorm:"column:user_data"`
	IsAktif   sql.NullBool   `gorm:"column:is_aktif"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
