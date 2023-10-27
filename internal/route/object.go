package route

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/golang-boilerplate-one/internal/app"
)

func ObjectRoute(a *fiber.App, handler *app.ServiceServer, enforcer *casbin.Enforcer) {
	route := a.Group("/object")
	// route.Get("/", handler.ListExample)
	route.Post("/", handler.CreateObject)
	route.Get("/:bucket/:id/view", handler.ViewObject)
	route.Get("/:bucket/:id/download", handler.DownloadObject)
	route.Get("/:id", handler.DetailObject)
	// route.Put("/:id", handler.PutExample)
	// route.Patch("/:id", handler.PatchExample)
	// route.Delete("/:id", handler.DeleteExample)
	// route.Delete("/:id", handler.DeleteExample)
}
