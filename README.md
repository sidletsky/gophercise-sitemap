## sitemap

Simple sitemap generator for the given url
```
sitemap [URL] [flags]
```

### Examples

Generating sitemap into ./sitemap.xml
```
sitemap https://duckduckgo.com/
```
Generating sitemap into ./my_own_name.xml
```
sitemap https://duckduckgo.com/ -f my_own_name.xml
```

### Options

```
  -f, --file string   name of an output file (default "sitemap.xml")
  -h, --help          help for sitemap
```

