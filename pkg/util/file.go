package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/riyan-eng/golang-boilerplate-one/env"
)

type fileStruct struct{}

func NewFile() *fileStruct {
	return &fileStruct{}
}

type fileMeta struct {
	Nama     string
	Tipe     string
	Size     int64
	MimeType string
	Path     string
	UUID     string
	Url      string
}

func (f *fileStruct) SaveLocal(c *fiber.Ctx, file *multipart.FileHeader, bucket string) (data fileMeta) {
	data.Nama = file.Filename
	b, err := file.Open()
	PanicIfNeeded(err)
	d, err := io.ReadAll(b)
	PanicIfNeeded(err)
	data.MimeType = http.DetectContentType(d)
	ext := filepath.Ext(file.Filename)
	data.Tipe = ext[1:]
	data.Size = file.Size / 1000

	data.UUID = uuid.NewString()
	filename := fmt.Sprintf(`%s.%s`, data.UUID, data.Tipe)
	data.Path = fmt.Sprintf(`./media/object/%s`, filename)

	data.Url = fmt.Sprintf(`%s/object/%s/%s/view`, env.SERVER_HOST_BE, bucket, data.UUID)
	go func() {
		c.SaveFile(file, data.Path)
	}()
	return
}

func (f *fileStruct) GetFileSizeString(size int64) (stringSize string) {
	switch len(strconv.Itoa(int(size))) {
	case 4:
		stringSize = fmt.Sprintf(`%.2f MB`, float64(size)/1000)
		return
	default:
		stringSize = fmt.Sprintf(`%v KB`, size)
		return
	}
}
