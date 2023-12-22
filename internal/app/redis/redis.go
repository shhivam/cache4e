package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"shivamsinghal.me/caching4e/internal/app/config"
)

func NewConnectionPool(ctx context.Context) (*redis.Client, error) {
	config := config.GetRedisConfig()
	rdb := redis.NewClient(config)

	resp := rdb.Ping(ctx)
	_, err := resp.Result()
	if err != nil {
		log.Fatalf("error pinging redis", err)
		return nil, err
	}

	fmt.Println("returned redis conn")

	return rdb, nil
}
