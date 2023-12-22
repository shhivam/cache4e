package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "pong"})
}

func HealthcheckDB(c *gin.Context, redisDB *redis.Client, postgresDB *sql.DB) {
	isRedisOkay := true

	redisResp := redisDB.Ping(c)
	_, err := redisResp.Result()
	if err != nil {
		isRedisOkay = false
	}

	isPostgresOkay := true
	err = postgresDB.Ping()
	if err != nil {
		isPostgresOkay = false
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   strconv.FormatBool(isRedisOkay && isPostgresOkay),
		"redis":    strconv.FormatBool(isRedisOkay),
		"postgres": strconv.FormatBool(isPostgresOkay),
	})
}
