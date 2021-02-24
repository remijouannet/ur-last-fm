package main

import (
	"fmt"
)

type User struct {
	Name    string `pg:",pk"`
	Country string
	Body    *UserJson
}

func (u User) String() string {
	return fmt.Sprintf("User<%s %s>", u.Name, u.Country)
}

type UserJson struct {
	User struct {
		Playlists  string `json:"playlists"`
		Playcount  string `json:"playcount"`
		Gender     string `json:"gender"`
		Name       string `json:"name"`
		Subscriber string `json:"subscriber"`
		URL        string `json:"url"`
		Country    string `json:"country"`
		Image      []struct {
			Size string `json:"size"`
			Text string `json:"#text"`
		} `json:"image"`
		Registered struct {
			Unixtime string `json:"unixtime"`
			Text     int    `json:"#text"`
		} `json:"registered"`
		Type      string `json:"type"`
		Age       string `json:"age"`
		Bootstrap string `json:"bootstrap"`
		Realname  string `json:"realname"`
	} `json:"user"`
}

type Track struct {
	Uts    int64 `pg:",pk"`
	Artist string
	Album  string
	Name   string `pg:",pk"`
	Body   *TrackJson
}

func (t Track) String() string {
	return fmt.Sprintf("Track<%d %s %s %s>", t.Uts, t.Artist, t.Album, t.Name)
}

type RecenttracksJson struct {
	Recenttracks struct {
		Attr struct {
			Page       string `json:"page"`
			Total      string `json:"total"`
			User       string `json:"user"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
		} `json:"@attr"`
		Track []TrackJson
	}
}

type TrackJson struct {
	Album struct {
		Text string `json:"#text"`
		Mbid string `json:"mbid"`
	} `json:"album"`
	Artist struct {
		Text string `json:"#text"`
		Mbid string `json:"mbid"`
	} `json:"artist"`
	Date struct {
		Text string `json:"#text"`
		Uts  string `json:"uts"`
	} `json:"date"`
	Image []struct {
		Text string `json:"#text"`
		Size string `json:"size"`
	} `json:"image"`
	Mbid       string `json:"mbid"`
	Name       string `json:"name"`
	Streamable string `json:"streamable"`
	URL        string `json:"url"`
}
