package dtoservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/riyan-eng/golang-boilerplate-one/internal/datastruct"
)

type AuthenticationLoginReq struct {
	Email    string
	Password string
	Issuer   string
}

type AuthenticationLoginRes struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    *jwt.NumericDate
	Match        bool
}

type AuthenticationRegisterReq struct {
	UUIDUser       string
	UUIDUserData   string
	Email        string
	Password     string
	Nama         string
	NIK          string
	KodeRole     string
	NomorTelepon string
}

type AuthenticationRefreshTokenReq struct {
	RefreshToken string
	Issuer       string
}

type AuthenticationRefreshTokenRes struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    *jwt.NumericDate
}

type AuthenticationValidateResetTokenReq struct {
	ResetToken string
}

type AuthenticationRequestResetToken struct {
	Email  string
	Issuer string
}

type AuthenticationResetPasswordReq struct {
	ResetToken string
	Password   string
}

type AuthenticationLogoutReq struct {
	IDUser int
}

type AuthenticationMeReq struct {
	IDUser int
}

type AuthenticationMeRes struct {
	Data datastruct.AuthenticationMe
}
