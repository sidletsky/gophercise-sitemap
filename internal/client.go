package internal

import (
	"io/ioutil"
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

// GetPage makes an http GetPage request
func (client *Client) GetPage(url string) ([]byte, error) {
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
	defer response.Body.Close()
	htmlPage, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return htmlPage, nil
}
