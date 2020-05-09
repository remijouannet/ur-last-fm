package main

import (
	"flag"
	"fmt"
	database "github.com/remijouannet/ur-last-fm/db"
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
	db       *database.Database
	debug    bool
)

func main() {
	var config string

	flag.StringVar(&token, "token", "", "specify a token for the api")
	flag.StringVar(&secret, "secret", "", "specify a secret for the api")
	flag.StringVar(&username, "username", "", "specify a username")
	flag.StringVar(&password, "password", "", "specify a password")
	flag.StringVar(&conn, "conn", "", "specify a conn string to connect to Postgresql")
	flag.StringVar(&scrap, "scrap", "", "specify a scrap action")
	flag.StringVar(&config, "config", "", "config file")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	log.Init(false)

	configFile(config)

	log.Init(debug)

	db = database.DbInit(conn, debug)

	log.Info(fmt.Sprintf("DataBase name is %s\n", db.GetDatabaseName()))

	scrap_split := strings.Split(scrap, ":")
	scrap_params := strings.Split(scrap_split[1], ",")

	log.Info(fmt.Sprintf("scrap action %s\n", scrap_split[0]))
	log.Info(fmt.Sprintf("scrap params %s\n", scrap_split[1]))

	switch scrap_split[0] {
	case "getUserInfo":
		getUserInfo(scrap_params[0])
	case "getAllRecentTracks":
		getAllRecentTracks(scrap_params[0])
	default:
		log.Error(fmt.Sprintf("scrap action not found %s.\n", scrap_split[0]))
	}

	//authGetMobileSession(P{"username": username, "password": password})
	//getAllRecentTracks("hoodlums36")
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "0"})
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "1"})
}
