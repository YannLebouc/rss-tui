package ui

import "github.com/YannLebouc/rss-tui/internal/rss"

type ErrMsg struct {
	err error
}

func (e ErrMsg) Error() string {
	return e.err.Error()
}

type FeedsLoadedMsg struct {
	Feeds []rss.Feed
}
