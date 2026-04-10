package fetch

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/YannLebouc/rss-tui/internal/feeds"
	"jaytaylor.com/html2text"
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

func safeHtmlToText(input string) (output string) {
	output, err := html2text.FromString(input, html2text.Options{})
	if err != nil {
		return input
	}
	return output
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

	root := feeds.Root{}
	feed := feeds.Feed{}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return feeds.Feed{}, err
	}

	if err := xml.Unmarshal(body, &root); err != nil {
		return feeds.Feed{}, err
	}

	switch root.XMLName.Local {
	case "rss":
		rssFeed := feeds.RssFeed{}
		if err := xml.Unmarshal(body, &rssFeed); err != nil {
			return feeds.Feed{}, err
		}

		feed.Title = safeHtmlToText(rssFeed.Channel.Title)
		feed.Link = rssFeed.Channel.Link
		feed.Date = rssFeed.Channel.PubDate

		for _, rssItem := range rssFeed.Channel.Items {
			item := feeds.Item{}

			item.Title = safeHtmlToText(rssItem.Title)
			item.Content = safeHtmlToText(rssItem.Description)
			item.Link = rssItem.Link
			item.Date = rssItem.PubDate

			feed.Items = append(feed.Items, item)
		}

	case "feed":
		atomFeed := feeds.AtomFeed{}
		if err := xml.Unmarshal(body, &atomFeed); err != nil {
			return feeds.Feed{}, err
		}

		feed.Title = safeHtmlToText(atomFeed.Title)
		feed.Date = atomFeed.Updated
		for _, link := range atomFeed.Links {
			if link.Rel == "alternate" || link.Rel == "" {
				feed.Link = link.Href
				break
			}
		}

		for _, entry := range atomFeed.Entries {
			item := feeds.Item{}

			item.Title = safeHtmlToText(entry.Title)

			if entry.Content != "" {
				item.Content = safeHtmlToText(entry.Content)
			} else {
				item.Content = safeHtmlToText(entry.Summary)
			}

			item.Date = entry.Updated
			for _, link := range entry.Links {
				if link.Rel == "alternate" {
					item.Link = link.Href
					break
				}
			}

			feed.Items = append(feed.Items, item)
		}
	}
	return feed, nil
}
