package main

import (
	"encoding/json"
	"fmt"
	log "github.com/remijouannet/ur-last-fm/log"
	"strconv"
)

func getAllRecentTracks(user string) {
	var result map[string]interface{}
	var recenttracks map[string]interface{}
	var track map[string]interface{}
	var date map[string]interface{}
	var artist map[string]interface{}
	var album map[string]interface{}

	var body []byte

	body = userGetRecentTracks(P{"user": user, "limit": "200"})
	json.Unmarshal([]byte(body), &result)

	recenttracks = result["recenttracks"].(map[string]interface{})

	total, _ := strconv.Atoi(recenttracks["@attr"].(map[string]interface{})["totalPages"].(string))
	log.Info(fmt.Sprintf("totalPages: %d\n", total))

	for i := 1; i <= total; i++ {
		body = userGetRecentTracks(P{"user": user, "limit": "200", "page": strconv.Itoa(i)})
		json.Unmarshal([]byte(body), &result)
		recenttracks = result["recenttracks"].(map[string]interface{})

		for _, value := range recenttracks["track"].([]interface{}) {
			track = value.(map[string]interface{})

			if _, ok := track["@attr"]; ok {
				continue
			}

			date = track["date"].(map[string]interface{})
			artist = track["artist"].(map[string]interface{})
			album = track["album"].(map[string]interface{})

			log.Info(fmt.Sprintf("track : %s %s %s\n", date["#text"], artist["#text"], album["#text"]))
		}
	}
}
