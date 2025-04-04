package controllers

import (
	"apigo/app/dto"
	"apigo/pkg/utils"
	"apigo/platform/cache"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func RenewTokens(ctx *fiber.Ctx) error {
	renewBody := &dto.RenewTokensRequest{}

	if err := ctx.BodyParser(renewBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validator, err := utils.NewModelsValidator()
	if err != nil {
		return utils.WrapInternalServerError("UserSignIn", err, ctx)
	}

	if err := validator.Struct(renewBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	payload, err := utils.ExtractJWTPayolad(ctx)
	if err != nil {
		return utils.WrapInternalServerError("RefreshTokens", err, ctx)
	}

	redisConn, err := cache.OpenRedisConnection()
	if err != nil {
		return utils.WrapInternalServerError("RefreshTokens", err, ctx)
	}

	getCmd := redisConn.Get(context.Background(), payload.UserID.String())
	redisRefreshToken, err := getCmd.Result()
	if err == redis.Nil {
		return utils.WrapUnauthorized(ctx, "refresh token probably has expired.")
	} else if err != nil {
		return utils.WrapInternalServerError("RefreshTokens", err, ctx)
	}

	if renewBody.RefreshToken != redisRefreshToken {
		return utils.WrapUnauthorized(ctx, "refresh token is invalid.")
	}

	tokens, err := utils.GenerateNewTokens(payload.UserID.String())
	if err != nil {
		return utils.WrapInternalServerError("RenewTokens", err, ctx)
	}

	err = redisConn.Set(context.Background(), payload.UserID.String(), tokens.Refresh, tokens.RefreshExpiresIn).Err()
	if err != nil {
		return utils.WrapInternalServerError("RenewTokens", err, ctx)
	}

	return ctx.JSON(dto.RenewTokensResponse{
		Message: "tokens successfully renewed.",
		Tokens: dto.Tokens{
			Access:  tokens.Access,
			Refresh: tokens.Refresh,
		},
	})
}
