package service

import (
	"database/sql"
	"strconv"

	"github.com/blockloop/scan/v2"
	"github.com/riyan-eng/golang-boilerplate-one/config"
	"github.com/riyan-eng/golang-boilerplate-one/env"
	"github.com/riyan-eng/golang-boilerplate-one/internal/datastruct"
	dtorepository "github.com/riyan-eng/golang-boilerplate-one/internal/dto_repository"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/internal/model"
	"github.com/riyan-eng/golang-boilerplate-one/internal/repository"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

type AuthenticationService interface {
	RefreshToken(req dtoservice.AuthenticationRefreshTokenReq) (res dtoservice.AuthenticationRefreshTokenRes)
	ResetPassword(req dtoservice.AuthenticationResetPasswordReq)
	Login(req dtoservice.AuthenticationLoginReq) (res dtoservice.AuthenticationLoginRes)
	Register(req dtoservice.AuthenticationRegisterReq)
	Logout(req dtoservice.AuthenticationLogoutReq)
	Me(req dtoservice.AuthenticationMeReq) (res dtoservice.AuthenticationMeRes)
	RequestResetToken(req dtoservice.AuthenticationRequestResetToken)
	ValidateResetToken(req dtoservice.AuthenticationValidateResetTokenReq)
}

type authenticationService struct {
	dao repository.DAO
}

func NewAuthenticationService(dao repository.DAO) AuthenticationService {
	return &authenticationService{
		dao: dao,
	}
}

func (a *authenticationService) RefreshToken(req dtoservice.AuthenticationRefreshTokenReq) (res dtoservice.AuthenticationRefreshTokenRes) {
	claim, err := util.ParseToken(req.RefreshToken, env.JWT_SECRET_REFRESH)
	if err != nil {
		panic(util.BadRequest{
			Message: "Invalid refresh token.",
		})

	}
	if err := util.ValidateToken(claim, "refresh"); err != nil {
		panic(util.BadRequest{
			Message: "Invalid refresh token.",
		})
	}
	genJwt := util.GenerateJwt(claim.UserID, claim.RoleCode, req.Issuer)
	res.AccessToken = genJwt.AccessToken
	res.RefreshToken = genJwt.RefreshToken
	res.ExpiredAt = genJwt.ExpiredAt
	return
}

func (a *authenticationService) ResetPassword(req dtoservice.AuthenticationResetPasswordReq) {
	claim, err := util.ParseToken(req.ResetToken, env.JWT_SECRET_RESET)
	if err != nil {
		panic(util.BadRequest{
			Message: "Invalid token.",
		})

	}
	if err := util.ValidateToken(claim, "resetPwd"); err != nil {
		panic(util.BadRequest{
			Message: "Invalid token.",
		})
	}
	// change password
	password := util.GenerateHash(req.Password)
	tabelUser:= model.User{
		ID: util.StringNumToInt(claim.UserID),
		Password: sql.NullString{String: password, Valid: util.ValidIsNotBlankString(password)},
	}
	a.dao.NewAuthenticationQuery().ResetPassword(dtorepository.AuthenticationResetPasswordReq{
		TabelUser: tabelUser,
	})

	// delete token
	a.dao.NewAuthenticationQuery().Logout(dtorepository.AuthenticationLogoutReq{
		IDUser: util.StringNumToInt(claim.UserID),
	})
}

func (a *authenticationService) Login(req dtoservice.AuthenticationLoginReq) (res dtoservice.AuthenticationLoginRes) {
	sqlRows := a.dao.NewAuthenticationQuery().Login(dtorepository.AuthenticationLoginReq{
		Email: req.Email,
	})
	var user datastruct.AuthenticationLogin
	err := scan.Row(&user, sqlRows)
	if err == sql.ErrNoRows {
		return
	} else {
		util.PanicIfNeeded(err)
	}

	enforce := config.NewEnforcer()
	enforce.AddRoleForUser(strconv.Itoa(user.ID), user.KodeRole)

	ok := util.VerifyHash(user.Password, req.Password)
	if ok {
		if !user.IsAktif {
			util.PanicIfNeeded(util.BadRequest{
				Message: "User is not active.",
			})
		}
		res.Match = true
		genJwt := util.GenerateJwt(strconv.Itoa(user.ID), user.KodeRole, req.Issuer)
		res.AccessToken = genJwt.AccessToken
		res.RefreshToken = genJwt.RefreshToken
		res.ExpiredAt = genJwt.ExpiredAt
		return
	}
	return
}

func (a *authenticationService) Register(req dtoservice.AuthenticationRegisterReq) {
	password := util.GenerateHash(req.Password)
	tabelUser := model.User{
		UUID:     req.UUIDUser,
		Email:    sql.NullString{String: req.Email, Valid: util.ValidIsNotBlankString(req.Email)},
		Password: sql.NullString{String: password, Valid: util.ValidIsNotBlankString(password)},
		Role:     sql.NullString{String: req.KodeRole, Valid: util.ValidIsNotBlankString(req.KodeRole)},
		UserData: sql.NullString{String: req.UUIDUserData, Valid: util.ValidIsNotBlankString(req.UUIDUserData)},
		IsAktif:  sql.NullBool{Bool: true, Valid: true},
	}
	tableUserData := model.UserData{
		UUID:         req.UUIDUserData,
		Nama:         sql.NullString{String: req.Nama, Valid: util.ValidIsNotBlankString(req.Nama)},
		NIK:          sql.NullString{String: req.NIK, Valid: util.ValidIsNotBlankString(req.NIK)},
		NomorTelepon: sql.NullString{String: req.NomorTelepon, Valid: util.ValidIsNotBlankString(req.NomorTelepon)},
	}
	a.dao.NewAuthenticationQuery().Register(dtorepository.AuthenticationRegisterReq{
		TabelUser:     tabelUser,
		TabelUserData: tableUserData,
	})
}

func (a *authenticationService) Logout(req dtoservice.AuthenticationLogoutReq) {
	a.dao.NewAuthenticationQuery().Logout(dtorepository.AuthenticationLogoutReq{
		IDUser: req.IDUser,
	})
}

func (a *authenticationService) Me(req dtoservice.AuthenticationMeReq) (res dtoservice.AuthenticationMeRes) {
	sqlRows := a.dao.NewAuthenticationQuery().Me(dtorepository.AuthenticationMeReq{
		IDUser: req.IDUser,
	})

	err := scan.Row(&res.Data, sqlRows)
	util.PanicIfNeeded(err)
	return
}

func (a *authenticationService) RequestResetToken(req dtoservice.AuthenticationRequestResetToken) {
	var user datastruct.AuthenticationRequestResetToken
	sqlrows := a.dao.NewAuthenticationQuery().RequestResetToken(dtorepository.AuthenticationRequestResetTokenReq{
		Email: req.Email,
	})
	err := scan.Row(&user, sqlrows)
	if err == sql.ErrNoRows {
		panic(util.BadRequest{
			Message: "Email tidak terdaftar.",
		})
	} else {
		util.PanicIfNeeded(err)
	}

	genToken := util.GenerateTokenResetPassword(strconv.Itoa(user.ID), user.KodeRole, req.Issuer)
	go func() {
		sender := util.NewGmailSender("SIPENTA", env.SMTP_EMAIL, env.SMTP_PASSWORD)
		subject := "Reset Password Verification"
		content := util.NewTemplate().EmailResetPassword(genToken.ResetPwdToken, genToken.ExpiredAt)
		to := []string{user.Email}
		err = sender.SendEmail(subject, content, to, nil, nil, nil)

	}()

}

func (a *authenticationService) ValidateResetToken(req dtoservice.AuthenticationValidateResetTokenReq) {
	claim, err := util.ParseToken(req.ResetToken, env.JWT_SECRET_RESET)
	if err != nil {
		panic(util.BadRequest{
			Message: "Invalid token.",
		})

	}
	if err := util.ValidateToken(claim, "resetPwd"); err != nil {
		panic(util.BadRequest{
			Message: "Invalid token.",
		})
	}
}
