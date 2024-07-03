package redis

import (
	"github.com/go-redis/redis"
	"github.com/km1110/task-copilot-server/pkg/config"
)

func NewRedisConnector() (*redis.Client, error) {
	redisConf := config.NewConfig().RedisConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	_, err := client.Ping().Result()

	if err != nil {
		return nil, err
	}

	return client, nil
}
