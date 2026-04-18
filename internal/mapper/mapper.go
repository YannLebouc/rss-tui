package mapper

import (
	"github.com/YannLebouc/rss-tui/internal/feeds"
	"github.com/YannLebouc/rss-tui/internal/processor"
)

func RssFeedToGenericFeed(rssFeed feeds.RssFeed) feeds.Feed {
	feed := feeds.Feed{}

	feed.Title = processor.SafeHtmlToText(rssFeed.Channel.Title)
	feed.Link = rssFeed.Channel.Link
	feed.Date = rssFeed.Channel.PubDate

	for _, rssItem := range rssFeed.Channel.Items {
		item := feeds.Item{}

		item.Title = processor.SafeHtmlToText(rssItem.Title)
		item.Content = processor.SafeHtmlToText(rssItem.Description)
		item.Link = rssItem.Link
		item.Date = rssItem.PubDate

		feed.Items = append(feed.Items, item)
	}

	return feed
}

func AtomFeedToGenericFeed(atomFeed feeds.AtomFeed) feeds.Feed {

	feed := feeds.Feed{}

	feed.Title = processor.SafeHtmlToText(atomFeed.Title)
	feed.Date = atomFeed.Updated
	for _, link := range atomFeed.Links {
		if link.Rel == "alternate" || link.Rel == "" {
			feed.Link = link.Href
			break
		}
	}

	for _, entry := range atomFeed.Entries {
		item := feeds.Item{}

		item.Title = processor.SafeHtmlToText(entry.Title)

		if entry.Content != "" {
			item.Content = processor.SafeHtmlToText(entry.Content)
		} else {
			item.Content = processor.SafeHtmlToText(entry.Summary)
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

	return feed
}
