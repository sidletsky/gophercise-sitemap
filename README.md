## sitemap

Simple sitemap generator for the given url
```bash
sitemap [URL] [flags]
```

### CLI usage

Generating sitemap into ./sitemap.xml
```bash
sitemap https://duckduckgo.com/
```
Generating sitemap into ./my_own_name.xml
```bash
sitemap https://duckduckgo.com/ -f my_own_name.xml
```

#### Options

```
  -f, --file string   name of an output file (default "sitemap.xml")
  -h, --help          help for sitemap
```

#### Output
Example output for `sitemap https://duckduckgo.com/`
```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://duckduckgo.com/bang/newbang/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/HTTP_Secure/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/HTTP_referrer/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/bang/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/about/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/donations/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/params/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/press/assets/press/DuckDuckGo-Brand-Logo.zip/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/press/assets/press/DuckDuckGo-Photos.zip/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/about/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/IP_Address/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/hiring/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/assets/email/DuckDuckGo-Privacy-Weekly_sample.png/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/press/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/spread/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/traffic/spread/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/Electronic_Frontier_Foundation/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/bang/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/User_agent/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/HTTP_cookie/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/feedback/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/traffic/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/app/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/privacy/settings/</loc>
    </url>
    <url>
        <loc>https://duckduckgo.com/about/hiring/</loc>
    </url>
</urlset>
```

### Package usage
Basic example
```go
package main

import (
	"fmt"
	"log"

	"github.com/sidletsky/sitemap"
)

func main() {
	urls, err := sitemap.Parse(" https://duckduckgo.com/", nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range urls {
		fmt.Println(url.Loc)
	}
}
```

With custom http client (in this case it's Google's App Engine)
```go
package main

import (
	"fmt"
	"log"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"github.com/sidletsky/sitemap"
)

...
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	urls, err := sitemap.Parse(" https://duckduckgo.com/", client)
	if err != nil {
		log.Fatal(err)
	}
	for _, url := range urls {
		fmt.Println(url.Loc)
	}
...
```
### Related
- [sitemap protocol specifications at https://www.sitemaps.org/protocol.html](https://www.sitemaps.org/protocol.html) 