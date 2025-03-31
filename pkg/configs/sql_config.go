package configs

import (
	"time"
)

type SQLConfig struct {
	DBType          string        `env:"DB_TYPE" validate:"required,oneof=postgres mysql"`
	Host            string        `env:"DB_HOST" validate:"required"`
	Port            uint8         `env:"DB_PORT" validate:"required"`
	User            string        `env:"DB_USER" validate:"required"`
	Password        string        `env:"DB_PASSWORD" validate:"required"`
	DBName          string        `env:"DB_NAME" validate:"required"`
	SSLMode         string        `env:"DB_SSL_MODE" validate:"required,oneof=enable disable"`
	MaxConns        int           `env:"DB_MAX_CONNECTIONS" validate:"required"`
	MaxIdleConns    int           `env:"DB_MAX_IDLE_CONNECTIONS" validate:"required"`
	MaxConnLifetime time.Duration `env:"DB_MAX_LIFETIME_CONNECTIONS" validate:"required"`
}

var sqlConfigInstance *SQLConfig

// Get the stored config or load it if not loaded yet.
func GetSQLConfig() (*SQLConfig, error) {
	if sqlConfigInstance != nil {
		return sqlConfigInstance, nil
	}

	err := loadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	config := &SQLConfig{}

	err = scanConfig(config)
	if err != nil {
		return nil, err
	}

	sqlConfigInstance = config
	return config, nil
}

// Reload the config or load it if not loaded yet.
func ReloadSQLConfig() (*SQLConfig, error) {
	err := reloadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	config := &SQLConfig{}

	err = scanConfig(config)
	if err != nil {
		return nil, err
	}

	sqlConfigInstance = config
	return config, nil
}
