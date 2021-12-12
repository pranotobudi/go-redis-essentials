package config

import "os"

type DbRedisConnection struct {
	HostPort string
	Host     string
	Port     string
	Password string
}

func NewDbRedisConfig() DbRedisConnection {
	dbConfig := DbRedisConnection{
		HostPort: os.Getenv("DB_REDIS_HOSTPORT"),
		Password: os.Getenv("DB_REDIS_PASSWORD"),
		Host:     os.Getenv("DB_REDIS_HOST"),
		Port:     os.Getenv("DB_REDIS_PORT"),
	}

	return dbConfig
}
