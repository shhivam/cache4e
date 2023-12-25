package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
)

func ReadThroughCache(
	c *gin.Context,
	redisDB *redis.Client,
	cacheKey string,
	expiry time.Duration,
	callbackOnCacheMiss func() (interface{}, error),
	dataStruct interface{}) error {

	result, err := redisDB.Get(c, cacheKey).Result()
	if errors.Is(err, redis.Nil) { // cache miss
		data, err := callbackOnCacheMiss()
		if err != nil {
			return err
		}
		jsonData, err := json.Marshal(data)
		// TODO: This seems like an unnecessary unmarshal. How to do this better?
		json.Unmarshal(jsonData, &dataStruct)
		if err != nil {
			return err
		}
		_, err = redisDB.Set(c, cacheKey, jsonData, expiry).Result()
		if err != nil {
			fmt.Println("redis key set error", err)
			return err
		}
		return nil
	} else { // cache hit
		err := json.Unmarshal([]byte(result), &dataStruct)
		return err
	}
}
