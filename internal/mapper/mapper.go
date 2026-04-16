package mapper

import (
	"github.com/YannLebouc/rss-tui/internal/feeds"
	"jaytaylor.com/html2text"
)

func RssFeedToGenericFeed(rssFeed feeds.RssFeed) feeds.Feed {
	feed := feeds.Feed{}

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

	return feed
}

func AtomFeedToGenericFeed(atomFeed feeds.AtomFeed) feeds.Feed {

	feed := feeds.Feed{}

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

	return feed
}

func safeHtmlToText(input string) (output string) {
	output, err := html2text.FromString(input, html2text.Options{})
	if err != nil {
		return input
	}
	return output
}
