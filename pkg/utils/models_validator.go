package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var modelsValidatorInstance *validator.Validate

func initValidator() error {
	validate := validator.New()

	err := validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		_, err := uuid.Parse(fl.Field().String())
		return err != nil
	})
	if err != nil {
		return err
	}

	modelsValidatorInstance = validate
	return nil
}

func NewModelsValidator() (*validator.Validate, error) {
	if modelsValidatorInstance != nil {
		return modelsValidatorInstance, nil
	}

	err := initValidator()
	if err != nil {
		log.Printf("Couldn't initiate models validator. Reason: %v", err)
		return nil, err
	}

	return modelsValidatorInstance, nil
}
