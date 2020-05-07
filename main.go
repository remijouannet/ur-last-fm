package main

import (
	"flag"
	orm "github.com/remijouannet/ur-last-fm/orm"
	"log"
	"os"
)

var (
	logInfo  *log.Logger
	logDebug *log.Logger
	logError *log.Logger
	token    string
	secret   string
	username string
	password string
	debug    bool
)

func main() {
	var config string

	logInfo = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	logDebug = log.New(os.Stdout, "DEBUG:", log.Ldate|log.Ltime)
	logError = log.New(os.Stdout, "ERROR:", log.Ldate|log.Ltime)

	flag.StringVar(&token, "token", "", "specify a token for the api")
	flag.StringVar(&secret, "secret", "", "specify a secret for the api")
	flag.StringVar(&username, "username", "", "specify a username")
	flag.StringVar(&password, "password", "", "specify a password")
	flag.StringVar(&config, "config", "", "config file")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	configFile(config)

	orm.Init()

	authGetMobileSession(P{"username": username, "password": password})
	getAllRecentTracks("hoodlums36")

	//userGetInfo("hoodlums36")
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "0"})
	//userGetRecentTracks(P{"user" : "hoodlums36", "limit": "1", "extended": "1"})
}
