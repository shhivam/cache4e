package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"shivamsinghal.me/caching4e/internal/app/dto"
	"shivamsinghal.me/caching4e/internal/app/postgres"
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

