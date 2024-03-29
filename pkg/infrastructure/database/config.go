package database

import (
	"os"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string
	option   string
}

func newConfigMySQL() *config {
	return &config{
		host:     os.Getenv("MYSQL_HOST"),
		database: os.Getenv("MYSQL_DATABASE"),
		port:     os.Getenv("MYSQL_PORT"),
		driver:   os.Getenv("MYSQL_DRIVER"),
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		option:   os.Getenv("MYSQL_OPTION"),
	}
}
