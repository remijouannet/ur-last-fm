package db

import (
	"database/sql"
	"fmt"
	log "github.com/remijouannet/ur-last-fm/log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func DbInit(conn string) (db *sql.DB, debug bool) {
	log.Init(debug)

	db, err := sql.Open("pgx", conn)
	if err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	log.Info("Connection sucessful to Postgresql\n")

	return
}
