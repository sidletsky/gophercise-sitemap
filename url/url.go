package url

import (
	"fmt"
	"strings"
)

type Url struct {
	Loc string
}

func New(loc string) Url {
	return Url{Loc: loc}
}

func (u Url) String() string {
	return fmt.Sprintf(`    <url>
        <loc>%s</loc>
    </url>`, u.Loc)
}

func removeQueryString(url string) string {
	if index := strings.Index(url, "?"); index != -1 {
		url = url[:index]
	}
	return url
}

func removeHash(url string) string {
	if index := strings.Index(url, "#"); index != -1 {
		url = url[:index]
	}
	return url
}

func Clean(url, domain string) string {
	// relative link (e.g. /content)
	if url == "/" || url == "" {
		return domain
	}
	if strings.HasSuffix(domain, "/") {
		domain = domain[:len(domain)-1]
	}
	if strings.HasPrefix(url, "/") {
		url = domain + url
	}
	url = removeQueryString(url)
	url = removeHash(url)
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return url
}

func InTargetDomain(url, domain string) bool {
	return strings.HasPrefix(url, domain)
}

func CleanBase(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	url = removeQueryString(url)
	url = removeHash(url)
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return url
}
