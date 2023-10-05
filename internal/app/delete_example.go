package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Delete
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /exampe/{id}/ [delete]
// @Security ApiKeyAuth
func (s *ServiceServer) DeleteExample(c *fiber.Ctx) error {
	s.exampleService.Delete(dtoservice.DeleteExampleReq{
		ID: util.NewQuery().GetIDByUUID("example", c.Params("id")),
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_DELETE)
}
