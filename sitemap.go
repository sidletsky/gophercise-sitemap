package sitemap

import (
	"errors"
	"strings"

	"github.com/sidletsky/sitemap/internal"
)

func Parse(baseUrl string) (*Node, error) {
	client := internal.NewClient(nil, baseUrl)
	sitemap := Node{url: baseUrl}
	err := buildSitemap(&client, sitemap.url, &sitemap)
	if err != nil {
		panic(err)
		return nil, err
	}
	sitemap.Print("")
	return nil, nil
}

func buildSitemap(client *internal.Client, baseUrl string, node *Node) error {
	links, err := client.GetPageLinks(node.url)
	if err != nil {
		return err
	}
	for _, link := range links {
		cleanLink, err := cleanUrl(link.Href, baseUrl)
		if err == nil {
			node.addChild(cleanLink)
		}
	}
	for _, child := range node.children {
		if !child.root().contains(child.url) {
			err := buildSitemap(client, baseUrl, child)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func cleanUrl(url, baseUrl string) (string, error) {
	// not in our website
	if !strings.HasPrefix(url, baseUrl) {
		return "", errors.New("not in targeted domain")
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
