package utils

import (
	"apigo/pkg/configs"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTPayolad struct {
	UserID    uuid.UUID
	ExpiresAt int64
}

func ExtractJWTPayolad(ctx *fiber.Ctx) (*JWTPayolad, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token is invalid or claims conversion error.")
	}

	userId, err := uuid.Parse(claims["id"].(string))
	if err != nil {
		return nil, err
	}

	expiresAt := claims["id"].(int64)

	return &JWTPayolad{
		UserID:    userId,
		ExpiresAt: expiresAt,
	}, nil
}

func verifyToken(ctx *fiber.Ctx) (*jwt.Token, error) {
	token := extractToken(ctx)
	return jwt.Parse(token, jwtKeyFunc)
}

func extractToken(ctx *fiber.Ctx) string {
	split := strings.Split(ctx.Get("Authorization"), ".")

	if len(split) == 2 {
		return split[1]
	} else {
		return ""
	}
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	config, err := configs.GetJWTConfig()
	if err != nil {
		return nil, err
	}

	return []byte(config.SecretKey), nil
}
