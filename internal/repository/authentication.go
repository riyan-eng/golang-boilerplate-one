package repository

import (
	"context"
	"database/sql"
	"fmt"

	dtorepository "github.com/riyan-eng/golang-boilerplate-one/internal/dto_repository"
	"github.com/riyan-eng/golang-boilerplate-one/internal/model"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthenticationQuery interface {
	Login(req dtorepository.AuthenticationLoginReq) *sql.Rows
	Register(req dtorepository.AuthenticationRegisterReq)
	RequestResetToken(req dtorepository.AuthenticationRequestResetTokenReq) *sql.Rows
	ResetPassword(req dtorepository.AuthenticationResetPasswordReq)
	Logout(req dtorepository.AuthenticationLogoutReq)
	Me(req dtorepository.AuthenticationMeReq) *sql.Rows
}

type authenticationQuery struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
	cache  *redis.Client
}

func (a *authenticationQuery) Login(req dtorepository.AuthenticationLoginReq) (sqlrows *sql.Rows) {
	query := fmt.Sprintf(`
	select u.id, u.uuid, u.email, u."password", u."role", u.is_aktif from users u where u.email = '%v' and u.is_aktif = true limit 1
	`, req.Email)
	sqlrows, err := a.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return
}

func (a *authenticationQuery) Register(req dtorepository.AuthenticationRegisterReq) {
	err := a.gormDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&req.TabelUserData).Error; err != nil {
			return err
		}
		if err := tx.Create(&req.TabelUser).Error; err != nil {
			return err
		}
		return nil
	})
	util.PanicIfNeeded(err)
}

func (a *authenticationQuery) RequestResetToken(req dtorepository.AuthenticationRequestResetTokenReq) (sqlRows *sql.Rows) {
	query := fmt.Sprintf(`
	select u.id, u.uuid, u.email, u."role" from users u where u.email = '%v' and u.is_aktif = true limit 1
	`, req.Email)
	sqlRows, err := a.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return
}

func (a *authenticationQuery) ResetPassword(req dtorepository.AuthenticationResetPasswordReq) {
	err := a.gormDB.Model(&model.User{}).Select("password").Where("id = ?", req.TabelUser.ID).Updates(req.TabelUser).Error
	util.PanicIfNeeded(err)
}

func (a *authenticationQuery) Logout(req dtorepository.AuthenticationLogoutReq) {
	ctx := context.Background()
	err := a.cache.Del(ctx, fmt.Sprintf("token-%v", req.IDUser)).Err()
	util.PanicIfNeeded(err)
}

func (a *authenticationQuery) Me(req dtorepository.AuthenticationMeReq) (sqlrows *sql.Rows) {
	query := fmt.Sprintf(`
	select u.id, u.uuid, u.email, u."role", r.nama as nama_role, u.is_aktif, coalesce(ud.nama,'') as nama, 
	coalesce(ud.nik, '') as nik, coalesce(ud.nomor_telepon, '') as nomor_telepon 
	from users u 
	left join roles r on u."role" = r.kode
	left join user_datas ud on u.user_data = ud.uuid 
	where u.id = %v limit 1
	`, req.IDUser)
	sqlrows, err := a.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return
}
