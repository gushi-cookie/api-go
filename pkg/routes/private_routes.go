package routes

import (
	"apigo/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(app *fiber.App) {
	router := app.Group("/api")

	router.Post("/user/signout", controllers.UserSignOut)
}