package internal

import (
	"fmt"
	"io"
	"log"
	"os"
)

const header = `<?xml version="1.0" encoding="UTF-8"?>
	<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
const footer = "</urlset>"

func CreateFile(file string, data map[string]Url) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	write(f, header)
	for _, v := range data {
		write(f, v.String())
	}
	write(f, footer)
}

func write(w io.Writer, a string) {
	_, err := fmt.Fprintln(w, a)
	if err != nil {
		log.Println(err)
	}
}
