package fetch

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/YannLebouc/rss-tui/internal/feeds"
	"github.com/YannLebouc/rss-tui/internal/mapper"
	"github.com/YannLebouc/rss-tui/internal/parser"
)

type Fetcher struct {
	httpClient *http.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (f *Fetcher) Fetch(url string) (feeds.Feed, error) {
	response, err := f.httpClient.Get(url)
	if err != nil {
		return feeds.Feed{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return feeds.Feed{}, fmt.Errorf("unexpected status code %d for %s", response.StatusCode, url)
	}

	var root feeds.Root
	var feed feeds.Feed

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return feeds.Feed{}, err
	}

	if err := xml.Unmarshal(body, &root); err != nil {
		return feeds.Feed{}, err
	}

	switch root.XMLName.Local {
	case "rss":
		rssFeed, err := parser.ParseRssFeed(body)
		if err != nil {
			return feeds.Feed{}, err
		}
		feed = mapper.RssFeedToGenericFeed(rssFeed)

	case "feed":
		atomFeed, err := parser.ParseAtomFeed(body)
		if err != nil {
			return feeds.Feed{}, err
		}
		feed = mapper.AtomFeedToGenericFeed(atomFeed)
	}

	return feed, nil
}
