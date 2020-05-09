package main

import (
	"encoding/json"
	"fmt"
	"github.com/remijouannet/ur-last-fm/log"
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
	Scrap    string `json:"scrap"`
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

	err = json.Unmarshal([]byte(jsonFile), &config)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error reading the JSON: %s \n", err))
		return
	}

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
	if scrap == "" && config.Scrap != "" {
		scrap = config.Scrap
	}
	if !debug && config.Debug {
		debug = config.Debug
	}
}
