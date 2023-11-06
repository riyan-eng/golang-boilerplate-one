package app

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary     Detail
// @Tags       	Example
// @Accept		json
// @Produce		json
// @Param       id		path	string				true	"id"
// @Router      /example/{id}/ [get]
// @Security ApiKeyAuth
func (s *ServiceServer) TemplateExample(c *fiber.Ctx) error {
	service := s.exampleService.Template()
	c.Response().Header.Set("Content-Type", "application/octet-stream")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=template_example.xlsx")
	return service.File.Write(c)
}
