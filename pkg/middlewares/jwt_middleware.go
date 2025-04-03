package middlewares

import (
	"apigo/pkg/configs"
	"apigo/pkg/utils"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func PrepareJWTMiddleware() (fiber.Handler, error) {
	jwtConfig, err := configs.GetJWTConfig()
	if err != nil {
		return nil, err
	}

	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(jwtConfig.SecretKey)},
		ErrorHandler: errorHandler,
	}

	return jwtware.New(config), nil
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	if strings.Contains(err.Error(), "Missing or malformed JWT") {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "token is invalid.",
		})
	} else if strings.Contains(err.Error(), "token is expired") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token has expired.",
		})
	}

	return utils.WrapInternalServerError("JWTMiddleware", err, ctx)
}
