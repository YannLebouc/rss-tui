package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/YannLebouc/rss-tui/internal/rss"
)

type Mode int

const (
	FEEDS_LIST Mode = iota
	ARTICLES_LIST
	ARTICLE_DETAIL
)

type Model struct {
	feeds           []rss.Feed
	mode            Mode
	selectedFeed    int
	selectedArticle int
	cursor          int
	loading         bool
	error           error
}

func InitialModel() Model {
	return Model{
		loading: true,
	}
}

func (m Model) Init() tea.Cmd {
	return LoadFeeds
}
