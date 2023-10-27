package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Detail
// @Tags       	Object
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /object/{id}/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) DetailObject(c *fiber.Ctx) error {
	service := s.objectService.Detail(dtoservice.DetailObjectReq{
		ID: util.NewQuery().GetIDByUUID("objects", c.Params("id")),
	})
	return util.NewResponse(c).Success(service.Item, nil, util.MESSAGE_OK_READ)
}
