package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type fiberConfigParsable struct {
	ReadTimeout int    `env:"SERVER_READ_TIMEOUT" validate:"required"`
	AppName     string `env:"SERVER_APP_NAME" envDefault:"api-go"`
	Host        string `env:"SERVER_HOST" validate:"required"`
	Port        uint16 `env:"SERVER_PORT" validate:"required"`
}

// Extended fiber config.
type FiberConfig struct {
	*fiber.Config
	Host string
	Port uint16
}

var fiberConfigInstance *FiberConfig

// Get the stored config or load it if not loaded yet.
func GetFiberConfig() (*FiberConfig, error) {
	if fiberConfigInstance != nil {
		return fiberConfigInstance, nil
	}

	err := loadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	fiberConfigInstance, err = createFiberConfig()
	return fiberConfigInstance, err
}

// Reload the config or load it if not loaded yet.
func ReloadFiberConfig() (*FiberConfig, error) {
	err := reloadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	fiberConfigInstance, err = createFiberConfig()
	return fiberConfigInstance, err
}

func createFiberConfig() (*FiberConfig, error) {
	parsed := fiberConfigParsable{}
	err := scanConfig(&parsed)
	if err != nil {
		return nil, err
	}

	config := &FiberConfig{
		Config: &fiber.Config{},
	}

	config.ReadTimeout = time.Second * time.Duration(parsed.ReadTimeout)
	config.Host = parsed.Host
	config.Port = parsed.Port
	config.AppName = parsed.AppName
	// VVV hardcoded properties
	config.CaseSensitive = true
	config.StrictRouting = true

	return config, nil
}
