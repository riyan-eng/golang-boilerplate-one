package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	qrcode "github.com/skip2/go-qrcode"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /example/{id}/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) DownloadExample(c *fiber.Ctx) error {
	var imgByte []byte
	imgByte, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)
	// bytea
	util.PanicIfNeeded(err)
	c.Write(imgByte)
	return nil
}
