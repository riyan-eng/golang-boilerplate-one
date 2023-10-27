package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /example/{id}/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) DetailExample(c *fiber.Ctx) error {
	service := s.exampleService.Detail(dtoservice.DetailExampleReq{
		ID: util.NewQuery().GetIDByUUID("example", c.Params("id")),
	})
	return util.NewResponse(c).Success(service.Item, nil, util.MESSAGE_OK_READ)
}
