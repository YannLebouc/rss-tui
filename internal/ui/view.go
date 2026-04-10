package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

func renderFeeds(model Model) tea.View {
	if len(model.feeds) == 0 {
		return tea.NewView("No feeds available")
	}

	s := "Press enter to view the feed articles :\n\n"

	for i, feed := range model.feeds {
		cursor := " " // no cursor
		if model.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, feed.Title)
	}

	s += "\nPress q to quit.\n"

	styledString := lipgloss.NewStyle().Width(model.width).Render(s)
	return tea.NewView(styledString)
}

func renderArticles(model Model) tea.View {
	if len(model.feeds[model.selectedFeed].Items) == 0 {
		return tea.NewView("No articles available")
	}

	s := "Press ENTER to open the article, press ESC to go back to feeds list\n\n"

	articles := model.feeds[model.selectedFeed].Items

	for i, article := range articles {
		cursor := " " // no cursor
		if model.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, article.Title)
	}

	s += "\nPress q to quit.\n"

	styledString := lipgloss.NewStyle().Width(model.width).Render(s)
	return tea.NewView(styledString)
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
		return tea.NewView(m.viewport.View())
	}

	return tea.NewView("RSS-TUI")
}
