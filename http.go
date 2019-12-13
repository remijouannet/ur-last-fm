package main

import (
    "io/ioutil"
    "net/http"
	"net/url"
    "encoding/json"
    "encoding/hex"
    "bytes"
    "crypto/md5"
    "sort"
    "strings"
)

const (
	UriApiSecBase  = "https://ws.audioscrobbler.com/2.0/"
	UriBrowserBase = "https://www.last.fm/api/auth/"
)

func MD5Hash(urlParams url.Values) string {
    var text string;
    keys := make([]string, 0, len(urlParams))

    for k := range urlParams {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    for _, k := range keys {
        text = text + k + urlParams[k][0]
        if debug {
            logDebug.Printf("Param: %s %s\n", k, urlParams[k][0])
        }
    }

    text = text + secret

    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func httpGet(params url.Values) (body []byte){
    var JSON bytes.Buffer

	params.Add("format", "json")

    if debug {
        for key, param := range params {
            logDebug.Printf("Param: %s %s\n", key, param[0])
        }
    }

	p := params.Encode()
    uri := UriApiSecBase + "?" + p

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

    if debug {
        logDebug.Printf("Status Code : %v\n", res.StatusCode)
        logDebug.Printf("Status : %s\n", res.Status)
    }

	body, err = ioutil.ReadAll(res.Body)

    if debug {
        json.Indent(&JSON, body, "", "\t")
        logDebug.Printf("Body : %d\n", len(JSON.Bytes()))
        logDebug.Printf("Body : %s\n", string(JSON.Bytes()))
    }

	return
}

func httpPost(params url.Values) (body []byte){
    var JSON bytes.Buffer

    params.Add("api_sig", MD5Hash(params))
	params.Add("format", "json")

    uri := UriApiSecBase

	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, strings.NewReader(params.Encode()))
	if err != nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

    if debug {
        logDebug.Printf("Status Code : %v\n", res.StatusCode)
        logDebug.Printf("Status : %s\n", res.Status)
    }

	body, err = ioutil.ReadAll(res.Body)

    if debug {
        json.Indent(&JSON, body, "", "\t")
        logDebug.Printf("Body : %d\n", len(JSON.Bytes()))
        logDebug.Printf("Body : %s\n", string(JSON.Bytes()))
    }

	return
}
