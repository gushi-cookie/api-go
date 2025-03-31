package database

import (
	"apigo/app/queries"
	"apigo/pkg/configs"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Queries struct {
	*queries.UserQueries
}

func OpenDBConnection() (*Queries, error) {
	var (
		db  *sqlx.DB
		err error
	)

	config, err := configs.GetSQLConfig()
	if err != nil {
		return nil, err
	}

	switch config.DBType {
	case "postgres":
		db, err = openPostgresConnection()
	case "mysql":
		db, err = openMySQLConnection()
	default:
		return nil, fmt.Errorf("error, database type '%s' not supported!", config.DBType)
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		&queries.UserQueries{DB: db},
	}, nil
}
