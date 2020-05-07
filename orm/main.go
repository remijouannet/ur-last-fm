package orm

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Init() {
	var log1 *log.Logger

	log1 = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	log1.Print("Test1\n")

	db, err := sql.Open("pgx", "user=lastfm password=password host=127.0.0.1 port=8081 database=lastfm sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

    err = db.Ping()
    if err != nil {
        panic(err)
    }

	log1.Print("Test success\n")
}
