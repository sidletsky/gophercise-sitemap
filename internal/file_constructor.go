package internal

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sidletsky/sitemap/url"
)

const header = `<?xml version="1.0" encoding="UTF-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
const footer = "</urlset>"

func CreateFile(file string, data map[string]url.Url) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	write(f, data)
}

func write(w io.Writer, data map[string]url.Url) {
	writeLine(w, header)
	for _, v := range data {
		writeLine(w, v.String())
	}
	writeLine(w, footer)
}

func writeLine(w io.Writer, a string) {
	_, err := fmt.Fprintln(w, a)
	if err != nil {
		log.Println(err)
	}
}
