package fetch

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/YannLebouc/rss-tui/internal/rss"
)

type Fetcher struct {
	httpClient *http.Client
}

func NewFetcher() *Fetcher {
	return &Fetcher{
		httpClient: &http.Client{},
	}
}

func (f *Fetcher) Fetch(url string) (rss.Feed, error) {
	response, err := f.httpClient.Get(url)
	if err != nil {
		return rss.Feed{}, err
	}

	defer response.Body.Close()

	if !(response.StatusCode == 200) {
		return rss.Feed{}, fmt.Errorf("HTTP reponse code expected : 200, got %d while trying to fetch %s", response.StatusCode, url)
	}

	feed := rss.Feed{}
	decoder := xml.NewDecoder(response.Body)
	if err := decoder.Decode(&feed); err != nil {
		return rss.Feed{}, err
	}

	return feed, nil
}
