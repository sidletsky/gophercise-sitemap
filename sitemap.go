package sitemap

import (
	"log"
	"net/http"

	"github.com/sidletsky/sitemap/internal"
	"github.com/sidletsky/sitemap/url"
)

type UrlMap map[string]url.Url

var httpClient *internal.HttpClient

func Parse(siteUrl string, client *http.Client) (UrlMap, error) {
	baseUrl := url.CleanBase(siteUrl)
	var err error
	httpClient, err = internal.NewClient(baseUrl, client)
	if err != nil {
		log.Fatal(err)
	}
	sitemap, err := buildSitemap(baseUrl)
	if err != nil {
		return nil, err
	}
	return sitemap, nil
}

func buildSitemap(baseUrl string) (UrlMap, error) {
	urls := make(UrlMap)
	urls[baseUrl] = url.Url{Loc: baseUrl}
	urls, err := buildSitemapRecursively(baseUrl, baseUrl, urls)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func buildSitemapRecursively(pageUrl, baseUrl string, urls UrlMap) (UrlMap, error) {
	links, err := httpClient.GetPageLinks(pageUrl)
	if err != nil {
		return nil, err
	}
	for _, link := range links {
		cleanLink := url.Clean(link.Loc, pageUrl)
		if _, ok := urls[cleanLink]; !ok && url.InTargetDomain(cleanLink, baseUrl) {
			urls[cleanLink] = url.Url{Loc: cleanLink}
			linkUrls, err := buildSitemapRecursively(cleanLink, baseUrl, urls)
			if err != nil {
				return nil, err
			}
			urls = unionMaps(urls, linkUrls)
		}
	}
	return urls, nil
}

func unionMaps(a, b UrlMap) UrlMap {
	for k, v := range b {
		a[k] = v
	}
	return a
}
