package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Put
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Param       body	body	httprequest.PutExample	true	"body"
// @Router      /exampe/{id}/ [put]
// @Security ApiKeyAuth
func (s *ServiceServer) PutExample(c *fiber.Ctx) error {
	// parse & validate
	body := new(httprequest.PutExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)

	s.exampleService.Put(dtoservice.PutExampleReq{
		ID:     util.NewQuery().GetIDByUUID("example", c.Params("id")),
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_UPDATE)
}
