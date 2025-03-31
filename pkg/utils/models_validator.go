package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var instance *validator.Validate

func initValidator() error {
	validate := validator.New()

	err := validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		_, err := uuid.Parse(fl.Field().String())
		return err != nil
	})
	if err != nil {
		return err
	}

	instance = validate
	return nil
}

func NewModelsValidator() (*validator.Validate, error) {
	if instance != nil {
		return instance, nil
	}

	err := initValidator()
	return instance, err
}
