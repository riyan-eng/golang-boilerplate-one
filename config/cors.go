package config

import "github.com/gofiber/fiber/v2/middleware/cors"

func NewCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		// AllowOrigins: "http://127.0.0.1:3000, https://gofiber.net, http://localhost:5173, https://dev-kanaltapem.sandboxindonesia.id",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		// AllowCredentials: true,
	}
}
