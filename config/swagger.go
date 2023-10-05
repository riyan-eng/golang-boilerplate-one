package config

import "github.com/gofiber/swagger"

func NewSwaggerConfig() swagger.Config {
	return swagger.Config{
		Title: "GOLANG",
	}
}
