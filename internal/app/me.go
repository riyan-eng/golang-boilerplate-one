package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Me
// @Tags       	Authentication
// @Produce		json
// @Router		/auth/me/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) Me(c *fiber.Ctx) error {
	idUser := util.StringNumToInt(c.Locals("user_id").(string))

	service := s.authService.Me(dtoservice.AuthenticationMeReq{
		IDUser: idUser,
	})
	return util.NewResponse(c).Success(service.Data, nil, util.MESSAGE_OK_READ)
}
