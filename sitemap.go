package gophercise_sitemap

import (
	"fmt"
	"github.com/sidletsky/gophercise-sitemap/internal"
)

type sitemap struct {
}

func Parse(url string) (*sitemap, error) {
	client := internal.NewClient(nil, url)
	page, err := client.GetPage(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(page))
	return nil, nil
}
