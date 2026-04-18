package service

import (
	"github.com/YannLebouc/rss-tui/internal/feeds"
	"github.com/YannLebouc/rss-tui/internal/fetch"
	"github.com/YannLebouc/rss-tui/internal/parser"
)

type FeedService struct {
	Fetcher *fetch.Fetcher
}

func (f *FeedService) GetFeed(url string) (feeds.Feed, error) {
	rawFeed, err := f.Fetcher.Fetch(url)
	if err != nil {
		return feeds.Feed{}, err
	}

	feed, err := parser.ParseToFeed(rawFeed)
	if err != nil {
		return feeds.Feed{}, err
	}

	return feed, nil
}
