package sitemap

import (
	"errors"
	"strings"

	"github.com/sidletsky/sitemap/internal"
)

func Parse(baseUrl string) (*Node, error) {
	client, err := internal.NewClient(nil, baseUrl)
	if err != nil {
		return nil, err
	}
	sitemap := Node{url: baseUrl}
	err = buildSitemap(client, sitemap.url, &sitemap)
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
		err := buildSitemap(client, baseUrl, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func cleanUrl(url, domain string) (string, error) {
	// relative link (e.g. /content)
	if url == "/" {
		return domain, nil
	}
	if strings.HasSuffix(domain, "/") {
		domain = domain[:len(domain)-1]
	}
	if strings.HasPrefix(url, "/") {
		url = domain + url
	}
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	// not in our website
	if !strings.HasPrefix(url, domain) {
		return "", errors.New("not in targeted domain")
	}
	return url, nil
}
