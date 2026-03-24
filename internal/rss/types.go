package rss

import "encoding/xml"

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Feed struct {
	URL     string   `xml:"-"`
	Tags    []string `xml:"-"`
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type FetchedFeed struct {
	Config FeedConfig
	Feed   Feed
}
