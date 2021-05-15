package main

import (
	"io"
	"testing"

	"github.com/remijouannet/ur-last-fm/log"
)

func TestConfigFile(t *testing.T) {

	log.Init(false, io.Discard)

	configFile("testdata/ur-last-fm.json")

	if token != "aaaaaaaaaaa" {
		t.Errorf("Failed to parse the config file")
	}
}
