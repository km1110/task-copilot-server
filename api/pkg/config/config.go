package config

import (
	"os"
	"time"
)

type config struct {
	dbInfo     *DBInfo
	cacheInfor *CacheInfo
}

type DBInfo struct {
	USER     string
	PASSWORD string
	DATABASE string
	HOST     string
	DRIVER   string
	PORT     string

	ENVIRONMENT string
}

type CacheInfo struct {
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

func NewConfig() *config {
	return &config{
		dbInfo:     &DBInfo{},
		cacheInfor: &CacheInfo{},
	}
}

func (c *config) DBConfig() DBInfo {
	c.dbInfo.USER = os.Getenv("POSTGRES_USER")
	c.dbInfo.PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	c.dbInfo.DATABASE = os.Getenv("POSTGRES_DB")
	c.dbInfo.HOST = os.Getenv("POSTGRES_HOST")
	c.dbInfo.DRIVER = os.Getenv("POSTGRES_DRIVER")
	c.dbInfo.PORT = os.Getenv("POSTGRES_PORT")
	c.dbInfo.ENVIRONMENT = os.Getenv("ENVIRONMENT")

	return *c.dbInfo
}

func (c *config) CacheConfig() CacheInfo {
	defaultExpirationStr := os.Getenv("CACHE_DEFAULT_EXPIRATION")
	cleanupIntervalStr := os.Getenv("CACHE_CLEANUP_INTERVAL")

	defaultExpiration, err := time.ParseDuration(defaultExpirationStr)
	if err != nil {
		defaultExpiration = 10 * time.Minute
	}

	cleanupInterval, err := time.ParseDuration(cleanupIntervalStr)
	if err != nil {
		cleanupInterval = 30 * time.Minute
	}

	c.cacheInfor.DefaultExpiration = defaultExpiration
	c.cacheInfor.CleanupInterval = cleanupInterval

	return *c.cacheInfor
}
