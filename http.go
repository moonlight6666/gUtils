package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"errors"
	"time"
)

// get
func HttpGet(apiURL string, data string) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	Url.RawQuery = data
	client := http.Client{
		Timeout: 120 * time.Second,
	}
	resp, err := client.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}


// post form
func HttpPostForm(url string, data string) ([]byte, error) {
	return HttpPost(url , "application/x-www-form-urlencoded", data )
}

// post json
func HttpPostJson(url string, data string) ([]byte, error) {
	return HttpPost(url , "application/json", data )
}

// post
func HttpPost(url string, contentType string, data string) ([]byte, error) {
	client := http.Client{
		Timeout: 120 * time.Second,
	}
	resp, err := client.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}
