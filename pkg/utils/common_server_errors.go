package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func WrapInternalServerError(handlerName string, err error, ctx *fiber.Ctx) error {
	log.Errorf("Handler '%s' has encountered an error. Reason: %v", handlerName, err)

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Something went wrong.",
	})
}

func WrapUnauthorized(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": message,
	})
}
