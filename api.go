package main

import (
	"github.com/remijouannet/ur-last-fm/log"
	"net/url"
)

type P map[string]string

func authGetMobileSession(params P) []byte {
	log.Debug("Method : authGetMobileSession\n")

	urlParams := url.Values{}
	urlParams.Add("method", "auth.getMobileSession")
	urlParams.Add("api_key", token)
	for key, param := range params {
		urlParams.Add(key, string(param))
	}

	return httpPost(urlParams)
}

func userGetInfo(params P) []byte {
	log.Debug("Method : userGetInfo\n")

	urlParams := url.Values{}
	urlParams.Add("method", "user.getInfo")
	urlParams.Add("api_key", token)
	for key, param := range params {
		urlParams.Add(key, string(param))
	}

	return httpGet(urlParams)
}

func userGetRecentTracks(params P) []byte {
	log.Debug("Method : userGetRecentTracks\n")

	if _, ok := params["limit"]; ok == false {
		params["limit"] = "5"
	}

	urlParams := url.Values{}
	urlParams.Add("method", "user.getRecentTracks")
	urlParams.Add("api_key", token)
	for key, param := range params {
		urlParams.Add(key, string(param))
	}

	return httpGet(urlParams)
}
