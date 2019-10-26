package main

import (
    "io/ioutil"
    "net/http"
	"net/url"
    "fmt"
    "flag"
)

const (
	UriApiSecBase  = "https://ws.audioscrobbler.com/2.0/"
	UriApiBase     = "http://ws.audioscrobbler.com/2.0/"
	UriBrowserBase = "https://www.last.fm/api/auth/"
)

func constructUrl(base string, params url.Values) (uri string) {
	//if ResponseFormat == "json" {
	//params.Add("format", ResponseFormat)
	//}
	p := params.Encode()
	uri = base + "?" + p
	return
}

func httpGet(method string, api_key string){
    fmt.Printf("%s\n", api_key)

    urlParams := url.Values{}
	urlParams.Add("method", method)
	urlParams.Add("api_key", api_key)

	uri := constructUrl(UriApiSecBase, urlParams)

	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}
    fmt.Printf("%v\n", res.StatusCode)
    fmt.Printf("%s\n", res.Status)
	body, err := ioutil.ReadAll(res.Body)
    fmt.Printf("%s\n", body)
	return
}

func main() {
    var token string
    flag.StringVar(&token, "token", "", "sepcify a token for the api")
    flag.Parse()
    fmt.Printf("token has value %s\n", token)
    httpGet("auth.getToken", token)
}
