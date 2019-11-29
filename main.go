package main

import (
    "io/ioutil"
    "net/http"
	"net/url"
    "fmt"
    "flag"
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

func MD5Hash(urlParams url.Values, api_secret string) string {
    var text string;
    keys := make([]string, 0, len(urlParams))

    for k := range urlParams {
        keys = append(keys, k)
    }

    sort.Strings(keys)

    for _, k := range keys {
        fmt.Printf("Params : %s\n", k)
        fmt.Printf("Params : %s\n", urlParams[k])
        text = text + k + urlParams[k][0]
    }

    text = text + api_secret

    hasher := md5.New()
    hasher.Write([]byte(text))
    fmt.Printf("TEXT : %s\n", text)
    fmt.Printf("MD5 : %s\n", hex.EncodeToString(hasher.Sum(nil)))
    return hex.EncodeToString(hasher.Sum(nil))
}

func httpGet(params url.Values) (body []byte){
    var JSON bytes.Buffer

	params.Add("format", "json")

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
    fmt.Printf("Status Code : %v\n", res.StatusCode)
    fmt.Printf("Status : %s\n", res.Status)
	body, err = ioutil.ReadAll(res.Body)

    json.Indent(&JSON, body, "", "\t")
    fmt.Printf("Body : %s\n", string(JSON.Bytes()))

	return
}

func httpPost(params url.Values, api_secret string) (body []byte){
    var JSON bytes.Buffer

    params.Add("api_sig", MD5Hash(params, api_secret))
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
    fmt.Printf("Status Code : %v\n", res.StatusCode)
    fmt.Printf("Status : %s\n", res.Status)
	body, err = ioutil.ReadAll(res.Body)

    json.Indent(&JSON, body, "", "\t")
    fmt.Printf("Body : %s\n", string(JSON.Bytes()))

	return
}


func authGetMobileSession(api_key string, api_secret string, username string, password string) (body []byte){
    urlParams := url.Values{}
	urlParams.Add("api_key", api_key)
	urlParams.Add("username", username)
	urlParams.Add("method", "auth.getMobileSession")
	urlParams.Add("password", password)

    body = httpPost(urlParams, api_secret)
    return
}

func userGetInfo(api_key string, user string) (body []byte){
    urlParams := url.Values{}
	urlParams.Add("method", "user.getInfo")
	urlParams.Add("api_key", api_key)
	urlParams.Add("user", user)

    body = httpGet(urlParams)
    return
}

func main() {
    var username string
    var password string
    var token string
    var secret string

    flag.StringVar(&token, "token", "", "specify a token for the api")
    flag.StringVar(&secret, "secret", "", "specify a secret for the api")
    flag.StringVar(&username, "username", "", "specify a username")
    flag.StringVar(&password, "password", "", "specify a password")
    flag.Parse()

    authGetMobileSession(token, secret, username, password)

    userGetInfo(token, "hoodlums36")
}
