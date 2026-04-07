package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"jaytaylor.com/html2text"
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

	return tea.NewView(s)
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

	return tea.NewView(s)
}

func renderArticleDetail(model Model) tea.View {
	article := model.feeds[model.selectedFeed].Items[model.selectedArticle]

	title, err := html2text.FromString(article.Title, html2text.Options{PrettyTables: true})
	content, err := html2text.FromString(article.Content, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}

	s := "Press ESC to go back to articles list\n\n"
	s += fmt.Sprintf("%s\n\n", title)
	s += fmt.Sprintf("%s\n\n", content)
	s += "\nPress q to quit.\n"

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
		return renderArticleDetail(m)
	}

	return tea.NewView("RSS-TUI")
}
