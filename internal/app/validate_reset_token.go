package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Validate Token Reset Password
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationValidateResetToken	true  "body"
// @Router		/auth/validate_reset_token/ [post]
func (s *ServiceServer) ValidateResetToken(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationValidateResetToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	s.authService.ValidateResetToken(dtoservice.AuthenticationValidateResetTokenReq{
		ResetToken: body.ResetToken,
	})
	return util.NewResponse(c).Success(nil, nil, "Valid")
}
