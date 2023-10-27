package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/golang-boilerplate-one/config"
	"github.com/riyan-eng/golang-boilerplate-one/internal/app"
	"github.com/riyan-eng/golang-boilerplate-one/internal/service"
)

func NewRoute(fiberApp *fiber.App,
	exampleService service.ExampleService,
	authenticationService service.AuthenticationService,
	objectService service.ObjectService,
) {
	allHandler := app.NewService(exampleService, authenticationService, objectService)
	enforcer := config.NewEnforcer()
	ExampleRoute(fiberApp, allHandler, enforcer)
	AuthenticationRoute(fiberApp, allHandler, enforcer)
	ObjectRoute(fiberApp, allHandler, enforcer)
}
