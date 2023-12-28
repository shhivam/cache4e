package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"shivamsinghal.me/caching4e/internal/app/cache"
	"shivamsinghal.me/caching4e/internal/app/dto"
	"shivamsinghal.me/caching4e/internal/app/postgres"
)

const userCacheKeyFormat string = "user:%s"

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
	cacheKey := fmt.Sprintf(userCacheKeyFormat, username)
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

func CreateUser(c *gin.Context, redisDB *redis.Client, postgresDB *sql.DB) {
	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    nil,
			"error":   "Couldn't parse the request body",
		})
		return
	}
	var user dto.User
	json.Unmarshal(reqBody, &user)

	user.Id = uuid.New().String()
	cacheKey := fmt.Sprintf(userCacheKeyFormat, user.Username)

	err = cache.WriteThroughCache(c, redisDB, postgresDB, user, cacheKey, postgres.InsertUser)

	if err != nil {
		// error handling can be more specific but I have kept it simple for now.
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
			"error":   "Error inserting user information in DB",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    user,
	})
}
