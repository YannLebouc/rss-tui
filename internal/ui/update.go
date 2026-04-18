package ui

import (
	"fmt"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if !m.ready {
			m.viewport = viewport.New(viewport.WithWidth(msg.Width), viewport.WithHeight(msg.Height))
			m.ready = true
		} else {
			m.viewport.SetWidth(msg.Width)
			m.viewport.SetHeight(msg.Height)
		}
		if m.mode == ARTICLE_DETAIL {
			m.viewport.SetContent(buildArticleContent(m))
		}

	case tea.KeyPressMsg:
		switch msg.String() {

		// Quitting the app
		case "ctrl+c", "q":
			return m, tea.Quit

		// Move up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// Move down
		case "down", "j":
			switch m.mode {
			case FEEDS_LIST:
				if m.cursor < len(m.feeds)-1 {
					m.cursor++
				}
			case ARTICLES_LIST:
				if m.cursor < len(m.feeds[m.selectedFeed].Items)-1 {
					m.cursor++
				}
			}

		// refreshing feeds
		case "r":
			m.loading = true
			return m, LoadFeeds(m.service)

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

				content := buildArticleContent(m)
				m.viewport.SetContent(content)
				m.viewport.GotoTop()
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

	if m.mode == ARTICLE_DETAIL {
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func buildArticleContent(m Model) string {
	article := m.feeds[m.selectedFeed].Items[m.selectedArticle]

	style := lipgloss.NewStyle().
		Width(m.contentWidth()).
		Padding(1, 2)

	s := fmt.Sprintf("%s\n\n%s\n\n[%d/%d]",
		article.Title,
		article.Content,
		m.selectedArticle+1,
		len(m.feeds[m.selectedFeed].Items),
	)

	return style.Render(s)
}

func (m Model) contentWidth() int {
	if m.viewport.Width() > 0 {
		return m.viewport.Width()
	}
	return 80
}
