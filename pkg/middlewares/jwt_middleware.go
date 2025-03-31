package middlewares

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ErrorHandler: errorHandler,
	}

	return jwtware.New(config)
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": true,
		"msg": err.Error(),
	})
}