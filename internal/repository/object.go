package repository

import (
	"database/sql"
	"fmt"

	dtorepository "github.com/riyan-eng/golang-boilerplate-one/internal/dto_repository"
	"github.com/riyan-eng/golang-boilerplate-one/internal/model"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"gorm.io/gorm"
)

type ObjectQuery interface {
	List(dtorepository.ListObjectReq) *sql.Rows
	Create(dtorepository.CreateObjectReq)
	Delete(dtorepository.DeleteObjectReq)
	Detail(dtorepository.DetailObjectReq) *sql.Rows
	Put(dtorepository.PutObjectReq)
	Patch(dtorepository.PatchObjectReq)
}

type objectQuery struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
}

func (t *objectQuery) List(req dtorepository.ListObjectReq) *sql.Rows {
	query := fmt.Sprintf(`select id, uuid, nama, coalesce(detail, '') as detail, created_at, updated_at, count(*) over() as total_rows from Object 
		where lower(nama) like lower('%%%v%%') order by nama %v limit %v offset %v`,
		req.Search, req.Order, req.Limit, req.Offset)
	rows, err := t.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (t *objectQuery) Create(req dtorepository.CreateObjectReq) {
	err := t.gormDB.Create(&req.Item).Error
	util.PanicIfNeeded(err)
}

func (t *objectQuery) Delete(req dtorepository.DeleteObjectReq) {
	err := t.gormDB.Delete(&model.Object{}, req.ID).Error
	util.PanicIfNeeded(err)
}

func (t *objectQuery) Detail(req dtorepository.DetailObjectReq) *sql.Rows {
	query := fmt.Sprintf(`
		select o.id, o.uuid, o.bucket, o.nama, o."size", o.mime_type, o.url, o.path, o.created_at, o.updated_at from objects o where o.id = %v
	`, req.ID)
	rows, err := t.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (t *objectQuery) Put(req dtorepository.PutObjectReq) {
	err := t.gormDB.Model(&model.Object{}).Select("nama", "detail").Where("id = ?", req.Item.ID).Updates(req.Item).Error
	util.PanicIfNeeded(err)
}

func (t *objectQuery) Patch(req dtorepository.PatchObjectReq) {
	err := t.gormDB.Model(&model.Object{}).Where("id = ?", req.Item.ID).Updates(req.Item).Error
	util.PanicIfNeeded(err)
}
