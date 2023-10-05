package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Refresh Token
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationRefreshToken	true  "body"
// @Router		/auth/refresh_token/ [post]
func (s *ServiceServer) RefreshToken(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationRefreshToken)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	service := s.authService.RefreshToken(dtoservice.AuthenticationRefreshTokenReq{
		RefreshToken: body.RefreshToken,
		Issuer:       string(c.Request().Host()),
	})
	data := fiber.Map{
		"access_token":  service.AccessToken,
		"refresh_token": service.RefreshToken,
		"expired_at":    service.ExpiredAt.Time,
	}
	return util.NewResponse(c).Success(data, nil, util.MESSAGE_OK_REFRESH)
}
