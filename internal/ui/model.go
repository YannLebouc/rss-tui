package ui

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/YannLebouc/rss-tui/internal/feeds"
)

type Mode int

const (
	FEEDS_LIST Mode = iota
	ARTICLES_LIST
	ARTICLE_DETAIL
)

type Model struct {
	width    int
	height   int
	ready    bool
	viewport viewport.Model

	feeds           []feeds.Feed
	mode            Mode
	selectedFeed    int
	selectedArticle int
	cursor          int

	loading bool
	error   error
}

func InitialModel() Model {
	return Model{
		loading: true,
		mode:    FEEDS_LIST,
	}
}

func (m Model) Init() tea.Cmd {
	return LoadFeeds
}
