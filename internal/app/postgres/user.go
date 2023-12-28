package postgres

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"log"
	"shivamsinghal.me/caching4e/internal/app/dto"
)

func ReadUser(c *gin.Context, postgresDB *sql.DB, username string) (*dto.User, error) {
	tx, err := postgresDB.BeginTx(c, nil)
	defer tx.Rollback()
	if err != nil {
		return nil, errors.New("couldn't initiate a database transaction")
	}
	row := tx.QueryRowContext(c, `SELECT * FROM users WHERE username = $1`, username)
	if err != nil {
		return nil, errors.New("couldn't SELECT rows from DB")
	}
	var user dto.User
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, errors.New("error querying user details")
	}

	return &user, nil
}

func InsertUser(c *gin.Context, postgresDB *sql.DB, data interface{}) error {
	user := data.(dto.User)
	tx, err := postgresDB.BeginTx(c, nil)
	defer tx.Rollback()
	if err != nil {
		return errors.New("couldn't initiate a database transaction")
	}

	insertStatement := `INSERT INTO users(id,username, password, email) VALUES ($1, $2, $3, $4)`
	result, err := tx.ExecContext(c, insertStatement, user.Id, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}

	numRowsAffected, err := result.RowsAffected()
	if numRowsAffected != 1 || err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
