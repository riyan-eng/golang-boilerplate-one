package datastruct

type AuthenticationLogin struct {
	ID       int    `db:"id"`
	UUID     string `db:"uuid"`
	Email    string `db:"email"`
	Password string `db:"password"`
	KodeRole string `db:"role"`
	IsAktif  bool   `db:"is_aktif"`
}

type AuthenticationRequestResetToken struct {
	ID       int    `db:"id"`
	UUID     string `db:"uuid"`
	Email    string `db:"email"`
	KodeRole string `db:"role"`
}

type AuthenticationMe struct {
	ID           int    `db:"id" json:"-"`
	UUID         string `db:"uuid" json:"id"`
	Email        string `db:"email" json:"email"`
	KodeRole     string `db:"role" json:"kode_role"`
	NamaRole     string `db:"nama_role" json:"nama_role"`
	Nama         string `db:"nama" json:"nama"`
	NIK          string `db:"nik" json:"nik"`
	NomorTelepon string `db:"nomor_telepon" json:"nomor_telepon"`
	IsAktif      bool   `db:"is_aktif" json:"is_aktif"`
}
