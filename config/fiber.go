package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/util"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: util.ErrorHandler,
	}
}
