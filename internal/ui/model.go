package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/YannLebouc/rss-tui/internal/rss"
)

type Model struct {
	feeds   []rss.Feed
	cursor  int
	loading bool
	error   error
}

func initialModel() Model {
	return Model{
		loading: true,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
