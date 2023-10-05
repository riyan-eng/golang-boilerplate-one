package app

import (
	"github.com/gofiber/fiber/v2"
	dtoservice "github.com/riyan-eng/golang-boilerplate-one/internal/dto_service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

// @Summary     Logout
// @Tags       	Authentication
// @Produce		json
// @Router		/auth/logout/ [delete]
// @Security ApiKeyAuth
func (s *ServiceServer) Logout(c *fiber.Ctx) error {
	idUser := util.StringNumToInt(c.Locals("user_id").(string))
	s.authService.Logout(dtoservice.AuthenticationLogoutReq{
		IDUser: idUser,
	})
	return util.NewResponse(c).Success(nil, nil, util.MESSAGE_OK_LOGOUT)
}
