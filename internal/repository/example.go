package repository

import (
	"database/sql"
	"fmt"

	dtorepository "github.com/riyan-eng/golang-boilerplate-one/internal/dto_repository"
	"github.com/riyan-eng/golang-boilerplate-one/internal/model"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"gorm.io/gorm"
)

type ExampleQuery interface {
	List(dtorepository.ListExampleReq) *sql.Rows
	Create(dtorepository.CreateExampleReq)
	Delete(dtorepository.DeleteExampleReq)
	Detail(dtorepository.DetailExampleReq) *sql.Rows
	Put(dtorepository.PutExampleReq)
	Patch(dtorepository.PatchExampleReq)
}

type exampleQuery struct {
	sqlDB  *sql.DB
	gormDB *gorm.DB
}

func (t *exampleQuery) List(req dtorepository.ListExampleReq) *sql.Rows {
	query := fmt.Sprintf(`select id, uuid, nama, coalesce(detail, '') as detail, created_at, updated_at, count(*) over() as total_rows from example 
		where lower(nama) like lower('%%%v%%') order by nama %v limit %v offset %v`,
		req.Search, req.Order, req.Limit, req.Offset)
	rows, err := t.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (t *exampleQuery) Create(req dtorepository.CreateExampleReq) {
	err := t.gormDB.Create(&req.Item).Error
	util.PanicIfNeeded(err)
}

func (t *exampleQuery) Delete(req dtorepository.DeleteExampleReq) {
	err := t.gormDB.Delete(&model.Example{}, req.ID).Error
	util.PanicIfNeeded(err)
}

func (t *exampleQuery) Detail(req dtorepository.DetailExampleReq) *sql.Rows {
	query := fmt.Sprintf(`
		select id, uuid, nama, coalesce(detail, '') as detail, created_at, updated_at from example where id = '%v'
	`, req.ID)
	rows, err := t.sqlDB.Query(query)
	util.PanicIfNeeded(err)
	return rows
}

func (t *exampleQuery) Put(req dtorepository.PutExampleReq) {
	err := t.gormDB.Model(&model.Example{}).Select("nama", "detail").Where("id = ?", req.Item.ID).Updates(req.Item).Error
	util.PanicIfNeeded(err)
}

func (t *exampleQuery) Patch(req dtorepository.PatchExampleReq) {
	err := t.gormDB.Model(&model.Example{}).Where("id = ?", req.Item.ID).Updates(req.Item).Error
	util.PanicIfNeeded(err)
}
