package main

import (
    "path/filepath"
    "os/user"
    "io/ioutil"
    "encoding/json"
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
