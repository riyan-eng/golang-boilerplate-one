package datastruct

import "time"

type Object struct {
	ID         int       `db:"id" json:"-"`
	UUID       string    `db:"uuid" json:"id"`
	Bucket     string    `db:"bucket" json:"bucket"`
	Nama       string    `db:"nama" json:"nama"`
	Size       int64     `db:"size" json:"-"`
	SizeString string    `json:"size"`
	MimeType   string    `db:"mime_type" json:"mime_type"`
	Url        string    `db:"url" json:"url"`
	Path       string    `db:"path" json:"path"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	TotalRows  int       `db:"total_rows" json:"-"`
}
