package main

import (
	"encoding/json"
	"fmt"
	log "github.com/remijouannet/ur-last-fm/log"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type Config struct {
	Token    string `json:"token"`
	Secret   string `json:"secret"`
	Username string `json:"username"`
	Password string `json:"password"`
	Conn     string `json:"conn"`
	Debug    bool   `json:"debug"`
}

func configFile(file string) {
	var config Config

	if file == "" {
		usr, _ := user.Current()
		dir := usr.HomeDir
		file = filepath.Join(dir, ".ur-last-fm.json")
	}

	log.Info(fmt.Sprintf("Trying to open: %s \n", file))

	jsonFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error(fmt.Sprintf("Error opening the config file: %s \n", err))
		return
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
	if conn == "" && config.Conn != "" {
		conn = config.Conn
	}
	if debug == false && config.Debug != false {
		debug = config.Debug
	}
}
