package main

import (
	"fmt"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/remijouannet/ur-last-fm/log"
)

func initDb(conn string) {
	log.Info(fmt.Sprintf("init db connection\n"))
	opts, err := pg.ParseURL(conn)
	log.PanicIf(err)

	db = pg.Connect(opts)
	db.AddQueryHook(pgdebug.DebugHook{Verbose: true})

	log.Info(fmt.Sprintf("create schema User connection\n"))

	err = db.Model((*User)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})
	if err != nil {
		panic(err)
	}

	err = db.Model((*User)(nil)).CreateTable(nil)
	if err != nil {
		panic(err)
	}

	log.Info(fmt.Sprintf("create schema Track connection\n"))

	err = db.Model((*Track)(nil)).DropTable(&orm.DropTableOptions{IfExists: true})
	if err != nil {
		panic(err)
	}

	err = db.Model((*Track)(nil)).CreateTable(nil)
	if err != nil {
		panic(err)
	}
}

func closeDb() {
	log.Info(fmt.Sprintf("close db connection\n"))
	db.Close()
}
