package main

import (
    "encoding/json"
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

    body = userGetRecentTracks(P{"user" : user, "limit": "200"})
    json.Unmarshal([]byte(body), &result)

    recenttracks = result["recenttracks"].(map[string]interface{})

    total, _ := strconv.Atoi(recenttracks["@attr"].(map[string]interface{})["totalPages"].(string))
    logInfo.Printf("total: %d\n", total)

    for i := 1; i <= total; i++ {
        body = userGetRecentTracks(P{"user" : user, "limit": "200", "page": strconv.Itoa(i)})
        json.Unmarshal([]byte(body), &result)
        recenttracks = result["recenttracks"].(map[string]interface{})

        for _, value := range recenttracks["track"].([]interface{}) {
            track = value.(map[string]interface{})
            date = track["date"].(map[string]interface{})
            artist = track["artist"].(map[string]interface{})
            album = track["album"].(map[string]interface{})

            logInfo.Printf("track : %s %s %s\n", date["uts"], artist["#text"], album["#text"])
        }
    }
}
