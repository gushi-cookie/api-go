package configs

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" validate:"required"`
	Port     uint16 `env:"REDIS_PORT" validate:"required"`
	Password string `env:"REDIS_PASSWORD" validate:"required"`
	DBNumber int    `env:"REDIS_DB_NUMBER"`
}

var redisConfigInstance *RedisConfig

// Get the stored config or load it if not loaded yet.
func GetRedisConfig() (*RedisConfig, error) {
	if redisConfigInstance != nil {
		return redisConfigInstance, nil
	}

	err := loadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	config := &RedisConfig{}

	err = scanConfig(config)
	if err != nil {
		return nil, err
	}

	redisConfigInstance = config
	return config, nil
}

// Reload the config or load it if not loaded yet.
func ReloadRedisConfig() (*RedisConfig, error) {
	err := reloadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	config := &RedisConfig{}

	err = scanConfig(&config)
	if err != nil {
		return nil, err
	}

	redisConfigInstance = config
	return config, nil
}
