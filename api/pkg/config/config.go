package config

import "os"

type config struct {
	dbInfo *DBInfo
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

func NewConfig() *config {
	return &config{
		dbInfo: &DBInfo{},
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
