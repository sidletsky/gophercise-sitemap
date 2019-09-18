package internal

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	link "github.com/sidletsky/gophercise-link"
)

type Client struct {
	HttpClient *http.Client
}

// New returns a Client that that wraps operations.
func NewClient(httpClient *http.Client, baseUrl string) (*Client, error) {
	var client Client
	if httpClient == nil {
		client.HttpClient = http.DefaultClient
	} else {
		client.HttpClient = httpClient
	}
	if !ping(baseUrl) {
		return nil, errors.New("destination host unreachable")
	}
	return &client, nil
}

func ping(url string) bool {
	out, _ := exec.Command("ping", url, "-c 5", "-i 3", "-w 10").Output()
	return !strings.Contains(string(out), "Destination Host Unreachable")
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

func (client *Client) GetPageLinks(url string) ([]link.Link, error) {
	page, err := client.GetPage(url)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(page)
	links, err := link.Parse(reader)
	if err != nil {
		return nil, err
	}
	return links, nil
}
