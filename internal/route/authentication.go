package route

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/golang-boilerplate-one/internal/app"
	"github.com/riyan-eng/golang-boilerplate-one/internal/middleware"
)

func AuthenticationRoute(a *fiber.App, handler *app.ServiceServer, enforcer *casbin.Enforcer) {
	route := a.Group("/auth")
	// route.Post("/register", middleware.AuthorizeJwt(), middleware.PermitCasbin(enforcer), handler.Register)
	route.Post("/register", handler.Register)
	route.Post("/login", handler.Login)
	route.Post("/refresh_token", handler.RefreshToken)
	route.Delete("/logout", middleware.AuthorizeJwt(), handler.Logout)
	route.Post("/request_reset_token", handler.RequestResetToken)
	route.Post("/reset_password", handler.ResetPassword)
	route.Get("/me", middleware.AuthorizeJwt(), handler.Me)
	route.Post("/validate_reset_token", handler.ValidateResetToken)
}
