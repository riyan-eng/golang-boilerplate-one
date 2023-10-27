package dtoservice

import "github.com/riyan-eng/golang-boilerplate-one/internal/datastruct"

type CreateObjectReq struct {
	UUID     string
	Bukcet   string
	Nama     string
	Size     int64
	MimeType string
	Url      string
	Path     string
}

type ListObjectReq struct {
	Search string
	Limit  int
	Offset int
	Order  string
}

type ListObjectRes struct {
	Items []datastruct.Object
	Total int
}

type DetailObjectReq struct {
	ID int
}

type DetailObjectRes struct {
	Item datastruct.Object
}

type DeleteObjectReq struct {
	ID int
}

type PutObjectReq struct {
	ID     int
	Nama   string
	Detail string
}

type PatchObjectReq struct {
	ID     int
	Nama   string
	Detail string
}
