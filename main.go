package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/riyan-eng/golang-boilerplate-one/config"
	"github.com/riyan-eng/golang-boilerplate-one/env"
	"github.com/riyan-eng/golang-boilerplate-one/infrastructure"
	"github.com/riyan-eng/golang-boilerplate-one/internal/repository"
	"github.com/riyan-eng/golang-boilerplate-one/internal/route"
	"github.com/riyan-eng/golang-boilerplate-one/internal/service"
	"github.com/riyan-eng/golang-boilerplate-one/pkg/docs"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}

	env.LoadEnvironmentFile()
	env.NewEnvironment()
	infrastructure.ConnectSqlDB()
	infrastructure.ConnectGormDB()
	infrastructure.ConnRedis()
	os.Setenv("TZ", env.SERVER_TIMEZONE)
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer access token here
func main() {
	// service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.GormDB, infrastructure.Redis, config.NewEnforcer())
	exampleService := service.NewExampleService(dao)
	authenticationService := service.NewAuthenticationService(dao)

	// swagger
	docs.SwaggerInfo.Title = "Golang Boilerplate One"
	fmt.Println(env.ENV)
	switch env.ENV {
	case "dev":
		docs.SwaggerInfo.Host = env.SERVER_HOST_BE
	default:
		docs.SwaggerInfo.Host = env.SERVER_HOST + ":" + env.SERVER_PORT
	}

	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	// fiber
	fiberApp := fiber.New(config.NewFiberConfig())
	fiberApp.Use(cors.New(config.NewCorsConfig()))
	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New())
	fiberApp.Get("/", func(c *fiber.Ctx) error { return c.SendString("Welcome to Golang Boilerplate One APIs") })
	fiberApp.Get("/metrics", monitor.New())
	fiberApp.Get("/docs/*", swagger.New(config.NewSwaggerConfig()))

	route.NewRoute(fiberApp, exampleService, authenticationService)
	fiberApp.Listen(env.SERVER_HOST + ":" + env.SERVER_PORT)
}
