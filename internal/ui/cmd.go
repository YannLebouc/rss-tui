package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/YannLebouc/rss-tui/internal/config"
	"github.com/YannLebouc/rss-tui/internal/fetch"
	"github.com/YannLebouc/rss-tui/internal/rss"
)

func LoadFeeds() tea.Msg {

	configPath, err := config.Path()
	if err != nil {
		return ErrMsg{err}
	}

	configLines, err := config.Load(configPath)
	if err != nil {
		return ErrMsg{err}
	}

	feeds := []rss.Feed{}
	fetcher := fetch.NewFetcher()
	for _, line := range configLines {
		feed, err := fetcher.Fetch(line.URL)
		if err != nil {
			return ErrMsg{err}
		}
		feeds = append(feeds, feed)
	}

	return FeedsLoadedMsg{
		feeds,
	}
}
