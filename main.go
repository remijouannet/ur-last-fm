package main

import (
    "flag"
    "log"
    "os"
)

var (
    logInfo *log.Logger
    logDebug *log.Logger
    token string
    secret string
    debug bool
)

func main() {
    var username string
    var password string

    flag.StringVar(&token, "token", "", "specify a token for the api")
    flag.StringVar(&secret, "secret", "", "specify a secret for the api")
    flag.StringVar(&username, "username", "", "specify a username")
    flag.StringVar(&password, "password", "", "specify a password")
    flag.BoolVar(&debug, "debug", false, "debug")
    flag.Parse()

    logInfo = log.New(os.Stdout, "INFO:", log.Ldate | log.Ltime)
    logDebug = log.New(os.Stdout, "DEBUG:", log.Ldate | log.Ltime)

    //authGetMobileSession(P{"username": username, "password": password})

    //userGetInfo("hoodlums36")

    userGetRecentTracks(P{"user" : "hoodlums36", "limit": "5"})
}
