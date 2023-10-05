package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	httprequest "github.com/riyan-eng/golang-boilerplate-one/pkg/http.request"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
	"github.com/riyan-eng/riyanisgood"
)

// @Summary     Login
// @Tags       	Authentication
// @Accept		json
// @Produce		json
// @Param       body	body  httprequest.AuthenticationLogin	true  "body"
// @Router		/auth/login/ [post]
func (s *ServiceServer) Login(c *fiber.Ctx) error {
	body := new(httprequest.AuthenticationLogin)
	err := c.BodyParser(&body)
	util.PanicIfNeeded(err)
	err, errors := riyanisgood.NewValidation().ValidateStruct(*body)
	if err != nil {
		return util.NewResponse(c).Error(errors, util.MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	service := s.authService.Login(dtoservice.AuthenticationLoginReq{
		Email:    body.Email,
		Password: body.Password,
		Issuer:   string(c.Request().Host()),
	})
	if !service.Match {
		return util.NewResponse(c).Error(nil, util.MESSAGE_FAILED_LOGIN, fiber.StatusBadRequest)
	}
	data := fiber.Map{
		"access_token":  service.AccessToken,
		"refresh_token": service.RefreshToken,
		"expired_at":    service.ExpiredAt.Time,
	}
	return util.NewResponse(c).Success(data, nil, util.MESSAGE_OK_LOGIN)
}
