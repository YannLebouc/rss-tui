package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/YannLebouc/rss-tui/internal/config"
	"github.com/YannLebouc/rss-tui/internal/feeds"
	"github.com/YannLebouc/rss-tui/internal/service"
)

func LoadFeeds(feedService *service.FeedService) tea.Cmd {
	return func() tea.Msg {
		configPath, err := config.Path()
		if err != nil {
			return ErrMsg{err}
		}

		configLines, err := config.Load(configPath)
		if err != nil {
			return ErrMsg{err}
		}

		var result = []feeds.Feed{}

		for _, line := range configLines {
			feed, err := feedService.GetFeed(line.URL)
			if err != nil {
				return ErrMsg{err}
			}
			result = append(result, feed)
		}

		return FeedsLoadedMsg{
			Feeds: result,
		}
	}
}
