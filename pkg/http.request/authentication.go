package httprequest

type AuthenticationRegister struct {
	Nama         string `json:"nama" valid:"required"`
	NIK          string `json:"nik" valid:"digits:16"`
	Email        string `json:"email" valid:"required;email"`
	NomorTelepon string `json:"nomor_telepon" valid:"digits_between:11,13"`
	Password     string `json:"password" valid:"required;min:6"`
}

type AuthenticationLogin struct {
	Email    string `json:"email" valid:"required;email"`
	Password string `json:"password" valid:"required"`
}

type AuthenticationRequestResetToken struct {
	Email string `json:"email" valid:"email"`
}

type AuthenticationResetPassword struct {
	ResetToken string `json:"token" valid:"required"`
	Password   string `json:"password" valid:"required"`
}

type AuthenticationValidateResetToken struct {
	ResetToken string `json:"token" valid:"required"`
}

type AuthenticationRefreshToken struct {
	RefreshToken string `json:"token" valid:"required"`
}
