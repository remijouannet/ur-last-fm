package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v10"
	log "github.com/remijouannet/ur-last-fm/log"
	"strconv"
)

func getUserInfo(db *pg.DB, user string) {
	var userjson *UserJson
	var body []byte
	var JSON bytes.Buffer

	body = userGetInfo(P{"username": user})
	json.Indent(&JSON, body, "", "\t")
	log.Info(fmt.Sprintf("Body : %s\n", JSON.String()))

	json.Unmarshal([]byte(body), &userjson)

	user1 := &User{
		Name:    userjson.User.Name,
		Country: userjson.User.Country,
		Body:    userjson,
	}
	_, err := db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}
}

func getAllRecentTracks(user string) {
	var result RecenttracksJson
	var body []byte
	var uts int

	body = userGetRecentTracks(P{"user": user, "limit": "200"})
	json.Unmarshal([]byte(body), &result)

	total, _ := strconv.Atoi(result.Recenttracks.Attr.TotalPages)
	log.Info(fmt.Sprintf("totalPages: %d\n", total))

	for i := 1; i <= total; i++ {
		body = userGetRecentTracks(P{"user": user, "limit": "200", "page": strconv.Itoa(i)})
		json.Unmarshal([]byte(body), &result)

		for _, track := range result.Recenttracks.Track {
			log.Info(fmt.Sprintf("track : %s %s %s %s\n", track.Date.Text, track.Artist.Text, track.Album.Text, track.Name))
			uts, _ = strconv.Atoi(track.Date.Uts)
			track1 := &Track{
				Uts:    int64(uts),
				Artist: track.Artist.Text,
				Album:  track.Album.Text,
				Name:   track.Name,
				Body:   &track,
			}
			_, err := db.Model(track1).Insert()
			if err != nil {
				panic(err)
			}
		}
	}
}
