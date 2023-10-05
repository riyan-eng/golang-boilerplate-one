package dtorepository

import "github.com/riyan-eng/golang-boilerplate-one/internal/model"

type AuthenticationRegisterReq struct {
	TabelUser     model.User
	TabelUserData model.UserData
}

type AuthenticationLoginReq struct {
	Email string
}

type AuthenticationLogoutReq struct {
	IDUser int
}

type AuthenticationRequestResetTokenReq struct {
	Email string
}

type AuthenticationResetPasswordReq struct {
	TabelUser model.User
}

type AuthenticationMeReq struct {
	IDUser int
}
