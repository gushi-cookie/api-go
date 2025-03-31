package main

import (
	"apigo/pkg/configs"
	"apigo/pkg/middlewares"
	"apigo/pkg/routes"
	"apigo/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := configs.GetFiberConfig()
	if err != nil {
		panic("to-do")
	}

	app := fiber.New(*config.Config)

	middlewares.FiberMiddleware(app)

	routes.PublicRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}