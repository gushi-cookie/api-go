package database

import (
	"apigo/pkg/configs"
	"apigo/pkg/utils"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func openPostgresConnection() (*sqlx.DB, error) {
	config, err := configs.GetSQLConfig()
	if err != nil {
		return nil, err
	}

	connURL, err := utils.ConnectionUrlBuilder("postgres")
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("pgx", connURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.MaxConnLifetime)

	return db, nil
}
