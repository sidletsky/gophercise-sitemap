package sitemap

import (
	"fmt"
	"strings"

	"github.com/sidletsky/sitemap/internal"
)

func Parse(baseUrl string) (*Node, error) {
	client, err := internal.NewClient(nil, baseUrl)
	if err != nil {
		return nil, err
	}
	sitemap, err := buildSitemap(client, baseUrl, nil)
	if err != nil {
		return nil, err
	}
	return sitemap, nil
}

func buildSitemap(client *internal.Client, baseUrl string, node *Node) (*Node, error) {
	if node == nil {
		node = &Node{url: baseUrl}
	}
	links, err := client.GetPageLinks(node.url)
	if err != nil {
		return node, err
	}
	for _, link := range links {
		cleanLink, err := cleanUrl(link.Href, baseUrl)
		if err == nil {
			node.addChild(cleanLink)
		}
	}
	for _, child := range node.children {
		node, err := buildSitemap(client, baseUrl, child)
		if err != nil {
			return node, err
		}
	}
	return node, nil
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
		return "", fmt.Errorf("not in target domain %s is not in %s", url, domain)
	}
	return url, nil
}
