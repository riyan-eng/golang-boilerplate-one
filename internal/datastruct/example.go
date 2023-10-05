package datastruct

import "time"

type Example struct {
	ID        int       `db:"id" json:"-"`
	UUID      string    `db:"uuid" json:"uuid"`
	Nama      string    `db:"nama" json:"nama"`
	Detail    string    `db:"detail" json:"detail"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	TotalRows int       `db:"total_rows" json:"-"`
}
