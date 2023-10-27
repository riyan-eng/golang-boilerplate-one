package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Download
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       bucket	path	string				true	"bucket"
// @Param       id		path	string				true	"id"
// @Router      /object/{bucket}/{id}/download [get]
// @Security ApiKeyAuth
func (s *ServiceServer) DownloadObject(c *fiber.Ctx) error {
	service := s.objectService.Detail(dtoservice.DetailObjectReq{
		ID: util.NewQuery().GetIDByUUID("objects", c.Params("id")),
	})
	c.Response().Header.SetContentType("application/octet-stream")
	c.Response().Header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", service.Item.Nama))
	return c.SendFile(service.Item.Path)
}
