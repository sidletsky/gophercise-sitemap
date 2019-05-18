package gophercise_sitemap

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	gophercise_link "github.com/sidletsky/gophercise-link"
	"github.com/sidletsky/gophercise-sitemap/internal"
)

func Parse(baseUrl string) (*Node, error) {
	client := internal.NewClient(nil, baseUrl)
	page, err := client.GetPage(baseUrl)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(page)
	links, err := gophercise_link.Parse(reader)
	if err != nil {
		return nil, err
	}
	sitemap := Node{url: baseUrl}
	for _, link := range links {
		cleanLink, err := clearUrl(link.Href, baseUrl)
		if err == nil {
			sitemap.addChild(cleanLink)
		}
	}
	for _, child := range sitemap.children {
		fmt.Println(child.url)
	}
	return nil, nil
}

func clearUrl(url, baseUrl string) (string, error) {
	// not in our website
	if !strings.HasPrefix(url, baseUrl) {
		return "", errors.New("Not in targeted domain")
	}
	ret := url
	// relative link (e.g. /content)
	if strings.HasPrefix(ret, "/") {
		ret = baseUrl + ret
	}
	if strings.HasSuffix(ret, "/") {
		ret = ret[:len(ret)-1]
	}
	return ret, nil
}
