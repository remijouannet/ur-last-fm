package db

import (
	"database/sql"
	"fmt"
	"github.com/remijouannet/ur-last-fm/log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Db           *sql.DB
	DatabaseName string
}

func DbInit(conn string, debug bool) *Database {
	log.Init(debug)

	NewDb, err := sql.Open("pgx", conn)
	if err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	err = NewDb.Ping()
	if err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	log.Info("Connection sucessful to Postgresql\n")

	return &Database{Db: NewDb}
}

func (db *Database) GetDatabaseName() string {
	if db.DatabaseName != "" {
		log.Info(fmt.Sprintf("DataBase name is %s\n", db.DatabaseName))
		return db.DatabaseName
	}

	rows, err := db.Db.Query("SELECT current_database();")
	if err != nil {
		log.Fatal(fmt.Sprint(err))
	}
	defer rows.Close()

	rows.Next()

	if err := rows.Scan(&db.DatabaseName); err != nil {
		log.Fatal(fmt.Sprint(err))
	}

	log.Debug(fmt.Sprintf("DataBase name is %s\n", db.DatabaseName))
	return db.DatabaseName
}
