package main

import (
	"fmt"
	"log"

	"github.com/YannLebouc/rss-tui/internal/config"
	"github.com/YannLebouc/rss-tui/internal/fetch"
)

func main() {
	configPath, err := config.Path()
	if err != nil {
		log.Fatal(err)
	}

	feeds, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	fetcher := fetch.NewFetcher()
	feed, err := fetcher.Fetch(feeds[0].URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(feed.Channel.Items[0])
}
