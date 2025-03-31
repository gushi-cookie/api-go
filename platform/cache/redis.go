package cache

import (
	"apigo/pkg/configs"
	"apigo/pkg/utils"

	"github.com/redis/go-redis/v9"
)

func OpenRedisConnection() (*redis.Client, error) {
	config, err := configs.GetRedisConfig()
	if err != nil {
		return nil, err
	}

	connUrl, err := utils.ConnectionUrlBuilder("redis")
	if err != nil {
		return nil, err
	}

	options := &redis.Options{
		Addr:     connUrl,
		Password: config.Password,
		DB:       config.DBNumber,
	}

	return redis.NewClient(options), nil
}
