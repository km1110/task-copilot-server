package cache

import (
	"github.com/km1110/task-copilot-server/pkg/config"
	"github.com/patrickmn/go-cache"
)

type IUserCache interface {
	Get(key string) string
	Set(key string, value interface{})
	Delete(key string)
}

type userCache struct {
	cache *cache.Cache
}

func NewUserCache() IUserCache {
	cacheConf := config.NewConfig().CacheConfig()
	return &userCache{
		cache: cache.New(cacheConf.DefaultExpiration, cacheConf.CleanupInterval),
	}
}

func (uc *userCache) Get(key string) string {
	value, found := uc.cache.Get(key)
	if !found {
		return ""
	}
	return value.(string)
}

func (uc *userCache) Set(key string, value interface{}) {
	uc.cache.Set(key, value, cache.DefaultExpiration)
}

func (uc *userCache) Delete(key string) {
	uc.cache.Delete(key)
}
