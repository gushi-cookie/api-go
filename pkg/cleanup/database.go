package cleanup

import (
	"apigo/platform/database"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

func CloseDBTransaction(handlerName string, tx *database.TxQueries, db *sqlx.DB, rollback *bool) {
	if rollback == nil {
		log.Errorf("argument 'rollback' is nil in handler '%s'.", handlerName)
		return
	}

	var err error

	if *rollback && tx != nil {
		err = tx.Rollback()
	}
	if err != nil {
		log.Errorf("couldn't rollback a transaction of handler '%s'. Reason: %v", handlerName, err)
	}

	if db != nil {
		err = db.Close()
	}
	if err != nil {
		log.Errorf("couldn't close a database connection of handler '%s'. Reason: %v", handlerName, err)
	}
}
