package main

import (
	"context"
	"github.com/joho/godotenv"
	"os"
	"shivamsinghal.me/caching4e/internal/app"
	"shivamsinghal.me/caching4e/internal/app/postgres"
	"shivamsinghal.me/caching4e/internal/app/redis"
)

func main() {
	err := godotenv.Load()
	ctx := context.Background()

	redisDB, err := redis.NewConnectionPool(ctx)
	if err != nil {
		os.Exit(1)
	}
	postgresDB, err := postgres.NewConnectionPool()
	if err != nil {
		os.Exit(1)
	}

	ginEngine := app.NewRouter(redisDB, postgresDB)

	ginEngine.Run()

}
