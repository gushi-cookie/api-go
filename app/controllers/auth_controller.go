package controllers

import (
	"apigo/app/dto"
	"apigo/app/models"
	"apigo/pkg/cleanup"
	"apigo/pkg/utils"
	"apigo/platform/cache"
	"apigo/platform/database"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserSignUp(ctx *fiber.Ctx) error {
	// 1. Parsing and validating the body
	signUp := &dto.SignUpRequest{}
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

	// 2. Opening a db transaction connection
	tx, db, err := database.OpenDBTransaction()
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	// 3. Preparing db cleanup
	shouldRollback := true
	defer cleanup.CloseDBTransaction("UserSignUp", tx, db, &shouldRollback)

	// 4. Checking if passed nickname and email are registered
	match, err := tx.HasUserByNicknameOrEmail(signUp.Nickname, signUp.Email)
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	} else if match {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "the email or nickname already registered.",
		})
	}

	// 5. Forming User and UserProfile valid instances
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

	profile := &models.UserProfile{}
	profile.UserId = user.ID
	profile.Nickname = signUp.Nickname
	profile.Bio = signUp.Bio

	if err := validator.Struct(profile); err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	// 6. Inserting models
	err = tx.CreateUser(user, profile)
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	err = tx.Commit()
	if err != nil {
		return utils.WrapInternalServerError("UserSignUp", err, ctx)
	}

	// 7. Making the response
	err = ctx.JSON(dto.SignUpResponse{
		Message: "user has been created.",
		User: dto.NewUser{
			ID:       user.ID,
			Nickname: profile.Nickname,
			Bio:      profile.Bio,
		},
	})
	if err != nil {
		return err
	}

	shouldRollback = false
	return nil
}

func UserSignIn(ctx *fiber.Ctx) error {
	signIn := &dto.SignInRequest{}

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

	return ctx.JSON(dto.SignInResponse{
		Message: "Ok",
		Tokens: dto.Tokens{
			Access:  tokens.Access,
			Refresh: tokens.Refresh,
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
