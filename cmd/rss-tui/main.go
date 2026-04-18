package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/YannLebouc/rss-tui/internal/fetch"
	"github.com/YannLebouc/rss-tui/internal/service"
	"github.com/YannLebouc/rss-tui/internal/ui"
)

func main() {

	fetcher := fetch.NewFetcher()
	feedService := service.FeedService{
		Fetcher: fetcher,
	}

	p := tea.NewProgram(ui.InitialModel(&feedService))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
