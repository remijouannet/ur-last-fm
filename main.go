package main

import (
	"flag"
	db "github.com/remijouannet/ur-last-fm/db"
	log "github.com/remijouannet/ur-last-fm/log"
)

var (
	token    string
	secret   string
	username string
	password string
	conn     string
	debug    bool
)

func main() {
	var config string

	flag.StringVar(&token, "token", "", "specify a token for the api")
	flag.StringVar(&secret, "secret", "", "specify a secret for the api")
	flag.StringVar(&username, "username", "", "specify a username")
	flag.StringVar(&password, "password", "", "specify a password")
	flag.StringVar(&conn, "conn", "", "specify a conn string to connect to Postgresql")
	flag.StringVar(&config, "config", "", "config file")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	log.Init(false)

	configFile(config)

	log.Init(debug)

	db.DbInit(conn)

	authGetMobileSession(P{"username": username, "password": password})
	//getAllRecentTracks("hoodlums36")

	userGetInfo(P{"username": "hoodlums36"})
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "0"})
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "1"})
}
