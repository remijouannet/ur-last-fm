package main

import (
    "flag"
    "log"
    "path/filepath"
    "os"
    "os/user"
    "io/ioutil"
    "encoding/json"
)

var (
    logInfo *log.Logger
    logDebug *log.Logger
    logError *log.Logger
    token string
    secret string
    username string
    password string
    debug bool
)

type Config struct {
  Token string `json:"token"`
  Secret string `json:"secret"`
  Username string `json:"username"`
  Password string `json:"password"`
  Debug bool `json:"debug"`
}

func configFile(file string){
    var config Config

    if file == "" {
        usr, _ := user.Current()
        dir := usr.HomeDir
        file = filepath.Join(dir, ".ur-last-fm.json")
    }

    logInfo.Printf("Trying to open: %s \n", file)

    jsonFile, err := ioutil.ReadFile(file)
    if err != nil {
        logError.Printf("Error opening the config file: %s \n", err)
    }
    json.Unmarshal([]byte(jsonFile), &config)

    if token == "" && config.Token != "" {
        token = config.Token
    }
    if secret == "" && config.Secret != "" {
        secret = config.Secret
    }
    if username == "" && config.Username != "" {
        username = config.Username
    }
    if password == "" && config.Password != "" {
        password = config.Password
    }
    if debug == false && config.Debug != false {
        debug = config.Debug
    }
}

func main() {
    var config string

    logInfo = log.New(os.Stdout, "INFO:", log.Ldate | log.Ltime)
    logDebug = log.New(os.Stdout, "DEBUG:", log.Ldate | log.Ltime)
    logError = log.New(os.Stdout, "ERROR:", log.Ldate | log.Ltime)

    flag.StringVar(&token, "token", "", "specify a token for the api")
    flag.StringVar(&secret, "secret", "", "specify a secret for the api")
    flag.StringVar(&username, "username", "", "specify a username")
    flag.StringVar(&password, "password", "", "specify a password")
    flag.StringVar(&config, "config", "", "config file")
    flag.BoolVar(&debug, "debug", false, "debug")
    flag.Parse()

    configFile(config)

    authGetMobileSession(P{"username": username, "password": password})
    getAllRecentTracks("hoodlums36")

    //userGetInfo("hoodlums36")
    //userGetRecentTracks(P{"user" : "hoodlums36", "limit": "5"})
}
