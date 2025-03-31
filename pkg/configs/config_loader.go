package configs

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var isConfigLoaded = false

// Load the config file and populate its values
// in the environments storage. Do nothing if
// the config is already loaded.
func loadDotEnvConfig() error {
	if isConfigLoaded { return nil }

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	isConfigLoaded = true
	return nil
}

// Force to load the config file again. Previous
// environment values will be overwritten.
func reloadDotEnvConfig() error {
	err := godotenv.Overload(".env")
	if err != nil {
		return err
	}

	isConfigLoaded = true
	return nil
}

// Parse and validate a config structure. It is supposed that
// corresponding environments are already loaded. Structure tags
// from these packages can be used:
//  "caarlos0/env/v11"
//  "go-playground/validator/v10"
func scanConfig(structure interface{}) error {
	err := env.Parse(&structure)
	if err != nil {
		return err
	}

	err = validator.New().Struct(&structure)
	if err != nil {
		return err
	}

	return nil
}