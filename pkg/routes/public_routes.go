package routes

import (
	"apigo/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	router := app.Group("/api")

	router.Post("/user/signup", controllers.UserSignUp)
	router.Post("/user/signin", controllers.UserSignIn)
}