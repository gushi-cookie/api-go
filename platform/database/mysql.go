package database

import (
	"apigo/pkg/configs"
	"apigo/pkg/utils"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func openMySQLConnection() (*sqlx.DB, error) {
	connURL, err := utils.ConnectionUrlBuilder("mysql")
	if err != nil {
		return nil, err
	}

	config, err := configs.GetSQLConfig()
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("mysql", connURL)
	db.SetMaxOpenConns(config.MaxConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.MaxConnLifetime)

	return db, nil
}