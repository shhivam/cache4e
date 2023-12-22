package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"shivamsinghal.me/caching4e/internal/app/config"
)

func NewConnectionPool() (*sql.DB, error) {
	pgConfig := config.GetPostgresConfig()
	pgConnectionStr := config.BuildPostgresConnString(pgConfig)
	db, err := sql.Open("postgres", pgConnectionStr)

	if err != nil {
		log.Fatalf("couldn't get connection to postgres")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Postgres DB ping failed", err)
		return nil, err
	}
	fmt.Println("returned pg conn")

	return db, nil
}
