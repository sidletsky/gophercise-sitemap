package sitemap

import (
	"fmt"
	"log"
	"strings"

	"github.com/sidletsky/sitemap/internal"
)

type UrlMap map[string]internal.Url

var httpClient internal.Client

func Parse(baseUrl string) (UrlMap, error) {
	client, err := internal.NewClient(nil, baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	httpClient = *client
	sitemap, err := buildSitemap(baseUrl)
	if err != nil {
		return nil, err
	}
	return sitemap, nil
}

func buildSitemap(baseUrl string) (UrlMap, error) {
	urls := make(UrlMap)
	urls[baseUrl] = internal.Url{Loc: baseUrl}
	urls, err := buildSitemapRecursively(baseUrl, urls)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func buildSitemapRecursively(baseUrl string, urls UrlMap) (UrlMap, error) {
	links, err := httpClient.GetPageLinks(baseUrl)
	if err != nil {
		return nil, err
	}
	for _, link := range links {
		cleanLink, err := cleanUrl(link.Href, baseUrl)
		if _, ok := urls[cleanLink]; !ok && err == nil {
			urls[cleanLink] = internal.Url{Loc: cleanLink}
			linkUrls, err := buildSitemapRecursively(cleanLink, urls)
			if err != nil {
				return nil, err
			}
			urls = join(urls, linkUrls)
		}
	}
	return urls, nil
}

func join(a, b UrlMap) UrlMap {
	for k, v := range b {
		a[k] = v
	}
	return a
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
