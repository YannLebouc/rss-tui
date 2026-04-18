package parser

import (
	"encoding/xml"
	"fmt"

	"github.com/YannLebouc/rss-tui/internal/feeds"
	"github.com/YannLebouc/rss-tui/internal/mapper"
)

func parseRssFeed(rawFeed []byte) (feeds.RssFeed, error) {
	rssFeed := feeds.RssFeed{}
	if err := xml.Unmarshal(rawFeed, &rssFeed); err != nil {
		return feeds.RssFeed{}, err
	}

	return rssFeed, nil
}

func parseAtomFeed(rawFeed []byte) (feeds.AtomFeed, error) {
	atomFeed := feeds.AtomFeed{}
	if err := xml.Unmarshal(rawFeed, &atomFeed); err != nil {
		return feeds.AtomFeed{}, err
	}

	return atomFeed, nil
}

func ParseToFeed(rawFeed []byte) (feeds.Feed, error) {

	var root feeds.Root

	if err := xml.Unmarshal(rawFeed, &root); err != nil {
		return feeds.Feed{}, err
	}

	switch root.XMLName.Local {
	case "rss":
		rssFeed, err := parseRssFeed(rawFeed)
		if err != nil {
			return feeds.Feed{}, err
		}
		return mapper.RssFeedToGenericFeed(rssFeed), nil

	case "feed":
		atomFeed, err := parseAtomFeed(rawFeed)
		if err != nil {
			return feeds.Feed{}, err
		}

		return mapper.AtomFeedToGenericFeed(atomFeed), nil

	default:
		return feeds.Feed{}, fmt.Errorf("unsupported feed type")
	}
}
