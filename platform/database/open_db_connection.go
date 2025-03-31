package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*sqlx.DB
}

func OpenDBConnection() (*Queries, error) {
	var (
		db *sqlx.DB
		err error
	)

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		db, err = openPostgresConnection()
	case "mysql":
		db, err = openMySQLConnection()
	default:
		return nil, fmt.Errorf("error, database type '%s' not supported!", dbType)
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		DB: db,
	}, nil
}