package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Create
// @Tags        Example
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.CreateExample	true  "body"
// @Router		/example/ [post]
// @Security ApiKeyAuth
func (s *ServiceServer) CreateExample(c *fiber.Ctx) error {
	body := new(httprequest.CreateExample)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}
	uuid := uuid.NewString()
	s.exampleService.Create(dtoservice.CreateExampleReq{
		UUID:   uuid,
		Nama:   body.Nama,
		Detail: body.Detail,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_CREATE)
}
