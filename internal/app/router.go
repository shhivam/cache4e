package app

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"shivamsinghal.me/caching4e/internal/app/controllers"
)

func NewRouter(redisDB *redis.Client, postgresDB *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck", controllers.Healthcheck)
	r.GET("/healthcheck-db", func(c *gin.Context) {
		controllers.HealthcheckDB(c, redisDB, postgresDB)
	})

	r.GET("/users/:username",
		func(c *gin.Context) {
			controllers.GetUser(c, postgresDB)
		})

	return r
}
