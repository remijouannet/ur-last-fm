package main

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/remijouannet/ur-last-fm/log"
)

type User struct {
	name       string
	url        string
	country    string
	registered int64
	//body map
}

func (u User) String() string {
	return fmt.Sprintf("User<%s %s %s %d>", u.name, u.url, u.country, u.registered)
}

func initDb(conn string) {
	log.Info(fmt.Sprintf("init db connection\n"))
	opts, err := pg.ParseURL(conn)
	log.PanicIf(err)

	db = pg.Connect(opts)

	err = createSchema(db)
	if err != nil {
		panic(err)
	}
}

func closeDb() {
	log.Info(fmt.Sprintf("close db connection\n"))
	db.Close()
}

func createSchema(db *pg.DB) error {
	log.Info(fmt.Sprintf("create schema connection\n"))
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}
