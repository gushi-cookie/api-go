package configs

type JWTConfig struct {
	SecretKey        string `env:"JWT_SECRET_KEY" validate:"required"`
	ExpiresInMinutes int    `env:"JWT_REFRESH_KEY" validate:"required"`
}

var jwtConfigInstance *JWTConfig

// Get the stored config or load it if not loaded yet.
func GetJWTConfig() (*JWTConfig, error) {
	if jwtConfigInstance != nil {
		return jwtConfigInstance, nil
	}

	err := loadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	config := &JWTConfig{}

	err = scanConfig(config)
	if err != nil {
		return nil, err
	}

	jwtConfigInstance = config
	return config, nil
}

// Reload the config or load it if not loaded yet.
func ReloadJWTConfig() (*JWTConfig, error) {
	err := reloadDotEnvConfig()
	if err != nil {
		return nil, err
	}

	config := &JWTConfig{}

	err = scanConfig(config)
	if err != nil {
		return nil, err
	}

	jwtConfigInstance = config
	return config, nil
}
