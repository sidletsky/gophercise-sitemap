package internal

import (
	"log"
	"net/http"
)

type Client struct {
	BaseUrl    string
	HttpClient *http.Client
}

// New returns a Client that that wraps operations.
func NewClient(httpClient *http.Client, baseUrl string) (client Client) {
	if httpClient == nil {
		client.HttpClient = http.DefaultClient
	} else {
		client.HttpClient = httpClient
	}
	client.BaseUrl = baseUrl
	return client
}

// Get makes an http Get request
func (client *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	response, err := client.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return response, nil
}
