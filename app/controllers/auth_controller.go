package controllers

import (
	"apigo/app/models"
	"apigo/pkg/utils"
	"apigo/platform/cache"
	"apigo/platform/database"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserSignUp(ctx *fiber.Ctx) error {
	signUp := &models.SignUp{}

	if err := ctx.BodyParser(signUp); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validator, err := utils.NewModelsValidator()
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	if err := validator.Struct(signUp); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	if exists, err := db.HasUserByEmail(signUp.Email); err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	} else if exists {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "the email already registered.",
		})
	}

	user := &models.User{}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Email = signUp.Email
	if hash, err := utils.HashPassword(signUp.Password); err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	} else {
		user.PassHash = hash
	}

	if err := validator.Struct(user); err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	err = db.CreateUser(user)
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	user.PassHash = ""
	return ctx.JSON(fiber.Map{
		"message": "User has been created.",
		"user":    user,
	})
}

func UserSignIn(ctx *fiber.Ctx) error {
	signIn := &models.SignIn{}

	if err := ctx.BodyParser(signIn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validator, err := utils.NewModelsValidator()
	if err != nil {
		return utils.WrapInternalServerError("UserSignIn", err, ctx)
	}

	if err := validator.Struct(signIn); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return utils.WrapInternalServerError("UserSignIn", err, ctx)
	}

	user, err := db.GetUserByEmail(signIn.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found or passed credentials are incorrect",
		})
	}

	if !utils.VerifyPassword(user.PassHash, signIn.Password) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found or passed credentials are incorrect",
		})
	}

	tokens, err := utils.GenerateNewTokens(user.ID.String())
	if err != nil {
		return utils.WrapInternalServerError("UserSignIn", err, ctx)
	}

	redisConn, err := cache.OpenRedisConnection()
	if err != nil {
		return utils.WrapInternalServerError("UserSignIn", err, ctx)
	}

	err = redisConn.Set(context.Background(), user.ID.String(), tokens.Refresh, tokens.RefreshExpiresIn).Err()
	if err != nil {
		return utils.WrapInternalServerError("UserSignIn", err, ctx)
	}

	return ctx.JSON(fiber.Map{
		"message": "OK",
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}

func UserSignOut(ctx *fiber.Ctx) error {
	payload, err := utils.ExtractJWTPayolad(ctx)
	if err != nil {
		return utils.WrapInternalServerError("UserSignOut", err, ctx)
	}

	redisConn, err := cache.OpenRedisConnection()
	if err != nil {
		return utils.WrapInternalServerError("UserSignOut", err, ctx)
	}

	err = redisConn.Del(context.Background(), payload.UserID.String()).Err()
	if err != nil {
		return utils.WrapInternalServerError("UserSignOut", err, ctx)
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
