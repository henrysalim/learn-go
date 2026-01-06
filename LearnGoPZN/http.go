package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	const BASE_URL string = "https://api.agify.io"
	params := url.Values{}
	params.Add("name", "Henry Salim")
	var fullURL string = BASE_URL + "?" + params.Encode()
	resp, err := http.Get(fullURL)

	if err != nil {
		errors.New("failed to make GET request: " + err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)         //should be '200 OK'
	fmt.Println(resp.Request.Method) //should be 'GET'

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errors.New("failed to read response body: " + err.Error())
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Request.Method)
	fmt.Println(string(body))
}
