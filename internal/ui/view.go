package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// The header
	if m.loading {
		s := "Loading..."
		return tea.NewView(s)
	}

	if m.error != nil {
		s := m.error.Error()
		return tea.NewView(s)
	}

	s := "Select the article you wish to read\n\n"

	// Iterate over our choices
	s += fmt.Sprintf("Articles from %s :\n", m.feeds[0].Channel.Title)
	for i, item := range m.feeds[0].Channel.Items {
		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, item.Title)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
