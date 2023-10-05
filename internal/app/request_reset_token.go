package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Request Reset Token
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationRequestResetToken	true  "body"
// @Router		/auth/request_reset_token/ [post]
func (s *ServiceServer) RequestResetToken(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationRequestResetToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	s.authService.RequestResetToken(dtoservice.AuthenticationRequestResetToken{
		Email:  body.Email,
		Issuer: string(c.Request().Host()),
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_REQUEST_TOKEN_RESET)
}
