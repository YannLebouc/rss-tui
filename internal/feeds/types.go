package feeds

import "encoding/xml"

type RssItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title   string    `xml:"title"`
	Link    string    `xml:"link"`
	PubDate string    `xml:"pubDate"`
	Items   []RssItem `xml:"item"`
}

type RssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type AtomLink struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Entry struct {
	Title   string     `xml:"title"`
	Content string     `xml:"content"`
	Summary string     `xml:"summary"`
	Links   []AtomLink `xml:"link"`
	Updated string     `xml:"updated"`
}

type AtomFeed struct {
	XMLName xml.Name   `xml:"feed"`
	Title   string     `xml:"title"`
	Links   []AtomLink `xml:"link"`
	Updated string     `xml:"updated"`
	Entries []Entry    `xml:"entry"`
}

type Item struct {
	Title   string
	Content string
	Link    string
	Date    string
}

type Feed struct {
	Title string
	Link  string
	Date  string
	Items []Item
}

type Root struct {
	XMLName xml.Name
}
