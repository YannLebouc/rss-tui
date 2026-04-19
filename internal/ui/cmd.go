package ui

import (
	"fmt"
	"time"

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

		type GetFeedResult struct {
			Feed feeds.Feed
			Err  error
		}

		var feeds = []feeds.Feed{}

		ch := make(chan GetFeedResult, len(configLines))

		for _, line := range configLines {
			go func(url string) {
				feed, err := feedService.GetFeed(url)
				ch <- GetFeedResult{Feed: feed, Err: err}
			}(line.URL)
		}

		timeout := time.After(time.Second * 5)

		for i := 0; i < len(configLines); i++ {
			select {
			case result := <-ch:
				if result.Err != nil {
					return ErrMsg{err}
				}
				feeds = append(feeds, result.Feed)
			case <-timeout:
				return ErrMsg{fmt.Errorf("timeout while fetching feeds")}
			}
		}

		return FeedsLoadedMsg{
			Feeds: feeds,
		}
	}
}
