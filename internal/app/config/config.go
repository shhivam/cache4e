package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

type PostgresConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func GetPostgresConfig() PostgresConfig {
	return PostgresConfig{
		host:     os.Getenv("PG_HOST"),
		port:     os.Getenv("PG_PORT"),
		user:     os.Getenv("PG_USER"),
		password: os.Getenv("PG_PASSWORD"),
		dbname:   os.Getenv("PG_DB_NAME"),
	}
}

func BuildPostgresConnString(config PostgresConfig) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		config.user,
		config.password,
		config.dbname,
		config.host,
		config.port,
	)
}

func GetRedisConfig() *redis.Options {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	fmt.Println("red host", host)
	return &redis.Options{
		Addr: host + ":" + port,
	}
}
