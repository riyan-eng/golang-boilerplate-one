package route

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/golang-boilerplate-one/internal/app"
)

func ExampleRoute(a *fiber.App, handler *app.ServiceServer, enforcer *casbin.Enforcer) {
	route := a.Group("/example")
	route.Get("/", handler.ListExample)
	route.Post("/", handler.CreateExample)
	route.Get("/:id", handler.DetailExample)
	route.Put("/:id", handler.PutExample)
	route.Patch("/:id", handler.PatchExample)
	route.Delete("/:id", handler.DeleteExample)
	route.Delete("/:id", handler.DeleteExample)
}
