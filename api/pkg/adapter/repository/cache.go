package repository

import "github.com/patrickmn/go-cache"

type ICacheRepository interface {
	Get(key string) string
	Set(key string, value interface{})
	Delete(key string)
}

type cacheRepository struct {
	cache *cache.Cache
}

func NewCacheRepository(cache *cache.Cache) ICacheRepository {
	return &cacheRepository{cache: cache}
}

func (cr *cacheRepository) Get(key string) string {
	value, found := cr.cache.Get(key)
	if !found {
		return ""
	}
	return value.(string)
}

func (cr *cacheRepository) Set(key string, value interface{}) {
	cr.cache.Set(key, value, cache.DefaultExpiration)
}

func (cr *cacheRepository) Delete(key string) {
	cr.cache.Delete(key)
}
