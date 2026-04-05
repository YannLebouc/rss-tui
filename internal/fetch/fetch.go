package fetch

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/YannLebouc/rss-tui/internal/feeds"
)

type Fetcher struct {
	httpClient *http.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		httpClient: &http.Client{},
	}
}

func (f *Fetcher) Fetch(url string) (feeds.Feed, error) {
	response, err := f.httpClient.Get(url)
	if err != nil {
		return feeds.Feed{}, err
	}

	defer response.Body.Close()

	if !(response.StatusCode == 200) {
		return feeds.Feed{}, fmt.Errorf("HTTP reponse code expected : 200, got %d while trying to fetch %s", response.StatusCode, url)
	}

	root := feeds.Root{}
	feed := feeds.Feed{}
	decoder := xml.NewDecoder(response.Body)
	if err := decoder.Decode(&root); err != nil {
		return feeds.Feed{}, err
	}

	switch root.XMLName.Local {
	case "rss":
		rssFeed := feeds.RssFeed{}
		if err := decoder.Decode(&rssFeed); err != nil {
			return feeds.Feed{}, err
		}

		feed.Title = rssFeed.Channel.Title
		feed.Link = rssFeed.Channel.Link
		feed.Date = rssFeed.Channel.PubDate

		for _, rssItem := range rssFeed.Channel.Items {
			item := feeds.Item{}

			item.Title = rssItem.Title
			item.Content = rssItem.Description
			item.Link = rssItem.Link
			item.Date = rssItem.PubDate

			feed.Items = append(feed.Items, item)
		}

	case "feed":
		atomFeed := feeds.AtomFeed{}
		if err := decoder.Decode(&atomFeed); err != nil {
			return feeds.Feed{}, err
		}

		feed.Title = atomFeed.Title
		feed.Date = atomFeed.Updated
		for _, link := range atomFeed.Links {
			if link.Rel == "self" {
				feed.Link = link.Href
				break
			}
		}

		for _, entry := range atomFeed.Entries {
			item := feeds.Item{}

			item.Title = entry.Title
			item.Content = entry.Content
			item.Date = entry.Updated
			for _, link := range entry.Links {
				if link.Rel == "self" {
					item.Link = link.Href
					break
				}
			}

			feed.Items = append(feed.Items, item)
		}
	}
	return feed, nil
}
