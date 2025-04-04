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

type TxQueries struct {
	*queries.UserTxQueries
}

func openConnection() (*sqlx.DB, error) {
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

	return db, nil
}

func OpenDBConnection() (*Queries, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		&queries.UserQueries{DB: db},
	}, nil
}

func OpenDBTransaction() (*TxQueries, *sqlx.DB, error) {
	db, err := openConnection()
	if err != nil {
		return nil, nil, err
	}

	tx, err := db.Beginx()
	if err != nil {
		return nil, nil, err
	}

	return &TxQueries{
		&queries.UserTxQueries{Tx: tx},
	}, db, nil
}
