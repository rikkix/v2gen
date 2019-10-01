package main

import (
	"io/ioutil"
	"net/http"
)

func GetContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

//TODO: Ping nodes
