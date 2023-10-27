package service

import (
	"database/sql"

	"github.com/blockloop/scan/v2"
	dtorepository "github.com/riyan-eng/golang-boilerplate-one/internal/dto_repository"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/internal/model"
	"github.com/riyan-eng/golang-boilerplate-one/internal/repository"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

type ObjectService interface {
	List(dtoservice.ListObjectReq) dtoservice.ListObjectRes
	Create(dtoservice.CreateObjectReq)
	Delete(dtoservice.DeleteObjectReq)
	Detail(dtoservice.DetailObjectReq) dtoservice.DetailObjectRes
	Put(dtoservice.PutObjectReq)
	Patch(dtoservice.PatchObjectReq)
}

type objectService struct {
	dao repository.DAO
}

func NewObjectService(dao repository.DAO) ObjectService {
	return &objectService{
		dao: dao,
	}
}

func (t *objectService) List(req dtoservice.ListObjectReq) (res dtoservice.ListObjectRes) {
	sqlrows := t.dao.NewObjectQuery().List(dtorepository.ListObjectReq{
		Search: req.Search,
		Limit:  req.Limit,
		Offset: req.Offset,
		Order:  req.Order,
	})
	err := scan.Rows(&res.Items, sqlrows)
	util.PanicIfNeeded(err)

	if len(res.Items) > 0 {
		res.Total = res.Items[0].TotalRows
	}
	return
}

func (t *objectService) Create(req dtoservice.CreateObjectReq) {
	item := model.Object{
		UUID:     req.UUID,
		Bucket:   sql.NullString{String: req.Bukcet, Valid: util.ValidIsNotBlankString(req.Bukcet)},
		Nama:     sql.NullString{String: req.Nama, Valid: util.ValidIsNotBlankString(req.Nama)},
		Size:     sql.NullInt64{Int64: req.Size, Valid: util.ValidIsNotZero(int(req.Size))},
		MimeType: sql.NullString{String: req.MimeType, Valid: util.ValidIsNotBlankString(req.MimeType)},
		Url:      sql.NullString{String: req.Url, Valid: util.ValidIsNotBlankString(req.Url)},
		Path:     sql.NullString{String: req.Path, Valid: util.ValidIsNotBlankString(req.Path)},
	}
	t.dao.NewObjectQuery().Create(dtorepository.CreateObjectReq{
		Item: item,
	})
}

func (t *objectService) Delete(req dtoservice.DeleteObjectReq) {
	t.dao.NewObjectQuery().Delete(dtorepository.DeleteObjectReq{
		ID: req.ID,
	})
}

func (t *objectService) Detail(req dtoservice.DetailObjectReq) (res dtoservice.DetailObjectRes) {
	sqlrows := t.dao.NewObjectQuery().Detail(dtorepository.DetailObjectReq{
		ID: req.ID,
	})
	err := scan.Row(&res.Item, sqlrows)
	util.PanicIfNeeded(err)
	res.Item.SizeString = util.NewFile().GetFileSizeString(res.Item.Size)
	return
}

func (t *objectService) Put(req dtoservice.PutObjectReq) {
	item := model.Object{
		ID: req.ID,
		// Nama: req.Nama,
	}
	t.dao.NewObjectQuery().Put(dtorepository.PutObjectReq{
		Item: item,
	})
}

func (t *objectService) Patch(req dtoservice.PatchObjectReq) {
	item := model.Object{
		ID: req.ID,
		// Nama: req.Nama,
	}
	t.dao.NewObjectQuery().Patch(dtorepository.PatchObjectReq{
		Item: item,
	})
}
