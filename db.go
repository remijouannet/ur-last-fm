package main

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/remijouannet/ur-last-fm/log"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

func initDb(conn string) {
	opts, err := pg.ParseURL(conn)
	log.PanicIf(err)

	db = pg.Connect(opts)

	var i int
	_, err = db.QueryOne(pg.Scan(&i), "SELECT 1")
	log.PanicIf(err)

	err = createSchema(db)
	log.PanicIf(err)
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		log.PanicIf(err)
	}
	return nil
}
