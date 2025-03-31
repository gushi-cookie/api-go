package utils

import (
	"apigo/pkg/configs"
	"fmt"
)

func ConnectionUrlBuilder(target string) (string, error) {
	var url string

	switch target {
	case "fiber":
		config, err := configs.GetFiberConfig()
		if err != nil {
			return "", err
		}

		url = fmt.Sprintf("%s:%d", config.Host, config.Port)
	case "mysql":
		config, err := configs.GetSQLConfig()
		if err != nil {
			return "", err
		}

		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
		)
	case "postgres":
		config, err := configs.GetSQLConfig()
		if err != nil {
			return "", err
		}

		url = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DBName,
			config.SSLMode,
		)
	case "redis":
		config, err := configs.GetRedisConfig()
		if err != nil {
			return "", err
		}

		url = fmt.Sprintf(
			"%s:%d",
			config.Host,
			config.Port,
		)
	default:
		return "", fmt.Errorf("connection target '%v' not supported", target)
	}

	return url, nil
}
