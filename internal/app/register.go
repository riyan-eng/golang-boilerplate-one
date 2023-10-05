package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Register
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationRegister	true  "body"
// @Router		/auth/register/ [post]
func (s *ServiceServer) Register(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationRegister)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	s.authService.Register(dtoservice.AuthenticationRegisterReq{
		UUIDUser:     uuid.NewString(),
		UUIDUserData: uuid.NewString(),
		Nama:         body.Nama,
		Email:        body.Email,
		Password:     body.Password,
		KodeRole:     "MASYARAKAT",
		NIK:          body.NIK,
		NomorTelepon: body.NomorTelepon,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_CREATE)
}
