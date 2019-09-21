package internal

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"golang.org/x/net/html"

	"github.com/sidletsky/sitemap/url"
)

type HttpClient struct {
	HttpClient *http.Client
}

var Client HttpClient

// NewClient returns a HttpClient that that wraps operations.
func NewClient(baseUrl string, httpClient *http.Client) (*HttpClient, error) {
	if httpClient == nil {
		Client.HttpClient = http.DefaultClient
	} else {
		Client.HttpClient = httpClient
	}
	if !ping(baseUrl) {
		return nil, errors.New("destination host unreachable")
	}
	return &Client, nil
}

func (client *HttpClient) GetPageLinks(url string) ([]url.Url, error) {
	page, err := client.getPage(url)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(page)
	links, err := getPageLinks(reader)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func ping(url string) bool {
	out, _ := exec.Command("ping", url, "-c 5", "-i 3", "-w 10").Output()
	return !strings.Contains(string(out), "Destination Host Unreachable")
}

// getPage makes an http get request
func (client *HttpClient) getPage(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	htmlPage, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return htmlPage, nil
}

// getPageLinks returns all links in html file
func getPageLinks(r io.Reader) ([]url.Url, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []url.Url
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// linkNodes returns list of nodes which are links
func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// buildLink returns a struct Url from an html.Node element
func buildLink(n *html.Node) url.Url {
	var ret url.Url
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Loc = attr.Val
			break
		}
	}
	return ret
}
