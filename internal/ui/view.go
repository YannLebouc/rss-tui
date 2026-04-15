package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func renderFeeds(m Model) tea.View {
	if len(m.feeds) == 0 {
		return tea.NewView("No feeds available")
	}

	s := "Select a feed (ENTER)\n\n"

	height := m.height - 4 // marge header/footer
	start, end := visibleRange(m.cursor, len(m.feeds), height)

	for i := start; i < end; i++ {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, m.feeds[i].Title)
	}

	s += "\nq: quit • ↑/↓: navigate • enter: select\n"

	return tea.NewView(s)
}

func renderArticles(m Model) tea.View {
	articles := m.feeds[m.selectedFeed].Items

	if len(articles) == 0 {
		return tea.NewView("No articles available")
	}

	s := "Articles (ENTER to open, ESC to go back)\n\n"

	height := m.height - 4
	start, end := visibleRange(m.cursor, len(articles), height)

	for i := start; i < end; i++ {
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		title := truncate(articles[i].Title, m.width-4)
		s += fmt.Sprintf("%s %s\n", cursor, title)
	}

	s += "\nq: quit • ↑/↓: navigate • enter: open • esc: back\n"

	return tea.NewView(s)
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

func visibleRange(cursor, total, height int) (start, end int) {
	if height <= 0 {
		return 0, total
	}

	half := height / 2

	start = cursor - half
	if start < 0 {
		start = 0
	}

	end = start + height
	if end > total {
		end = total
		start = max(0, end-height)
	}

	return start, end
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}
