package parser

import (
	"encoding/xml"

	"github.com/YannLebouc/rss-tui/internal/feeds"
)

func ParseRssFeed(responseBody []byte) (feeds.RssFeed, error) {
	rssFeed := feeds.RssFeed{}
	if err := xml.Unmarshal(responseBody, &rssFeed); err != nil {
		return feeds.RssFeed{}, err
	}

	return rssFeed, nil
}

func ParseAtomFeed(responseBody []byte) (feeds.AtomFeed, error) {
	atomFeed := feeds.AtomFeed{}
	if err := xml.Unmarshal(responseBody, &atomFeed); err != nil {
		return feeds.AtomFeed{}, err
	}

	return atomFeed, nil
}
