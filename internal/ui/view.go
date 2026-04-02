package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func renderFeeds(model Model) tea.View {
	s := "Select a feed by pressing enter to get it's articles :\n\n"

	for i, feed := range model.feeds {
		cursor := " " // no cursor
		if model.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, feed.Channel.Title)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}

func renderArticles(model Model) tea.View {
	return tea.NewView("")
}

func renderArticleDetail(model Model) tea.View {
	return tea.NewView("")
}

func (m Model) View() tea.View {

	if m.loading {
		s := "Loading..."
		return tea.NewView(s)
	}

	if m.error != nil {
		s := m.error.Error()
		return tea.NewView(s)
	}

	switch m.mode {
	case FEEDS_LIST:
		return renderFeeds(m)
	case ARTICLES_LIST:
		return renderArticles(m)
	case ARTICLE_DETAIL:
		return renderArticleDetail(m)
	}

	return tea.NewView("RSS-TUI")
}
