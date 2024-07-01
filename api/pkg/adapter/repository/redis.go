package repository

import "github.com/go-redis/redis"

type IRedisRepository interface {
	Get(key string) (string, error)
	Set(key string, value interface{})
	Delete(key string)
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *redisRepository {
	return &redisRepository{client: client}
}

func (rr *redisRepository) Get(key string) (string, error) {
	value, err := rr.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rr *redisRepository) Set(key string, value interface{}) {
	rr.client.Set(key, value, 0)
}

func (rr *redisRepository) Delete(key string) {
	rr.client.Del(key)
}
