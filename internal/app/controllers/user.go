package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"shivamsinghal.me/caching4e/internal/app/cache"
	"shivamsinghal.me/caching4e/internal/app/dto"
	"shivamsinghal.me/caching4e/internal/app/postgres"
	"time"
)

func GetUser(c *gin.Context, postgresDB *sql.DB) {
	username := c.Param("username")

	user, err := postgres.ReadUser(c, postgresDB, username)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    gin.H{},
			})
		log.Fatalf("%v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    user,
	})
}

func GetCachedUser(c *gin.Context, redisDB *redis.Client, postgresDB *sql.DB) {
	username := c.Param("username")
	cacheKey := fmt.Sprintf("user:%s", username)
	fmt.Println("username", username)
	callback := func() (interface{}, error) {
		fmt.Println("callback called")
		return postgres.ReadUser(c, postgresDB, username)
	}

	var user dto.User
	err := cache.ReadThroughCache(c, redisDB, cacheKey, 15*time.Second, callback, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    gin.H{},
		})
		log.Fatalf("%v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    user,
	})
}
