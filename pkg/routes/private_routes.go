package routes

import (
	"apigo/app/controllers"
	"apigo/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	router := app.Group("/api")

	jwtHandler, err := middlewares.PrepareJWTMiddleware()
	if err != nil {
		panic("to-do")
	}

	router.Post("/user/signout", jwtHandler, controllers.UserSignOut)
}
