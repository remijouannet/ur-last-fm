package main

import (
	"net/url"
)

type P map[string]string

func authGetMobileSession(params P) []byte{
    logInfo.Print("Method : authGetMobileSession\n")

    urlParams := url.Values{}
	urlParams.Add("method", "auth.getMobileSession")
	urlParams.Add("api_key", token)
    for key, param := range params {
        urlParams.Add(key, string(param[0]))
    }

    return httpPost(urlParams)
}

func userGetInfo(params P) []byte{
    logInfo.Print("Method : userGetInfo\n")

    urlParams := url.Values{}
	urlParams.Add("method", "user.getInfo")
	urlParams.Add("api_key", token)
    for key, param := range params {
        urlParams.Add(key, string(param[0]))
    }

    return httpGet(urlParams)
}

func userGetRecentTracks(params P) []byte{
    logInfo.Print("Method : userGetRecentTracks\n")

    if _, ok := params["limit"]; ok == false {
        params["limit"] = "5"
    }

    urlParams := url.Values{}
	urlParams.Add("method", "user.getRecentTracks")
	urlParams.Add("api_key", token)
    for key, param := range params {
        urlParams.Add(key, string(param[0]))
    }

    return httpGet(urlParams)
}
