package ui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		// Quitting the app
		case "ctrl+c", "q":
			return m, tea.Quit

		// Moving cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// Moving cursor down
		case "down", "j":
			if m.mode == FEEDS_LIST {
				if m.cursor < len(m.feeds)-1 {
					m.cursor++
				}
			}
			if m.mode == ARTICLES_LIST {
				if m.cursor < len(m.feeds[m.selectedFeed].Items)-1 {
					m.cursor++
				}
			}

		// refreshing feeds
		case "r":
			m.loading = true
			return m, LoadFeeds

		// Selecting a feed or an article
		case "enter":
			switch m.mode {
			case FEEDS_LIST:
				m.selectedFeed = m.cursor
				m.cursor = 0
				m.mode = ARTICLES_LIST
			case ARTICLES_LIST:
				m.selectedArticle = m.cursor
				m.mode = ARTICLE_DETAIL
			}

		// Going back to previous view
		case "esc":
			switch m.mode {
			case ARTICLES_LIST:
				m.mode = FEEDS_LIST
				m.cursor = 0
				m.selectedFeed = m.cursor
			case ARTICLE_DETAIL:
				m.mode = ARTICLES_LIST
				m.cursor = 0
				m.selectedArticle = m.cursor
			}
		}

	case FeedsLoadedMsg:
		m.feeds = msg.Feeds
		m.mode = FEEDS_LIST
		m.loading = false
		m.error = nil
		m.cursor = 0
		m.selectedArticle = m.cursor
		m.selectedFeed = m.cursor

	case ErrMsg:
		m.error = msg.err
		m.loading = false
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
