package dtoservice

import (
	"github.com/riyan-eng/golang-boilerplate-one/internal/datastruct"
	"github.com/xuri/excelize/v2"
)

type CreateExampleReq struct {
	UUID   string
	Nama   string
	Detail string
}

type ListExampleReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type ListExampleRes struct {
	Items []datastruct.Example
	Total int
}

type DetailExampleReq struct {
	ID int
}

type DetailExampleRes struct {
	Item datastruct.Example
}

type DeleteExampleReq struct {
	ID int
}

type PutExampleReq struct {
	ID     int
	Nama   string
	Detail string
}

type PatchExampleReq struct {
	ID     int
	Nama   string
	Detail string
}

type TemplateExampleRes struct {
	File *excelize.File
}
