package main

import (
	"flag"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/remijouannet/ur-last-fm/log"
	"strings"
)

var (
	token    string
	secret   string
	username string
	password string
	conn     string
	scrap    string
	conflict string
	db       *pg.DB
	debug    bool
)

func main() {
	var config string
	log.Init(false)

	flag.StringVar(&token, "token", "", "specify a token for the api")
	flag.StringVar(&secret, "secret", "", "specify a secret for the api")
	flag.StringVar(&username, "username", "", "specify a username")
	flag.StringVar(&password, "password", "", "specify a password")
	flag.StringVar(&conn, "conn", "", "specify a conn string to connect to Postgresql")
	flag.StringVar(&scrap, "scrap", "", "specify a scrap action")
	flag.StringVar(&conflict, "conflict", "stop", "specify an action on conflit in the database: continue, stop")
	flag.StringVar(&config, "config", "", "config file")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	configFile(config)
	log.Init(debug)

	defer closeDb()

	initDb(conn)

	scrap_split := strings.Split(scrap, ":")
	scrap_params := strings.Split(scrap_split[1], ",")
	log.Info(fmt.Sprintf("scrap action %s\n", scrap_split[0]))
	log.Info(fmt.Sprintf("scrap params %s\n", scrap_split[1]))

	switch scrap_split[0] {
	case "getUserInfo":
		getUserInfo(db, scrap_params[0], conflict)
	case "getAllRecentTracks":
		getAllRecentTracks(scrap_params[0], conflict)
	default:
		log.Error(fmt.Sprintf("scrap action not found %s.\n", scrap_split[0]))
	}

	//authGetMobileSession(P{"username": username, "password": password})
	//getAllRecentTracks("hoodlums36")
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "0"})
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "1"})
}
