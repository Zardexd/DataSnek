package main

import "github.com/gocolly/colly/v2"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
}
