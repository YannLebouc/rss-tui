package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type FeedConfig struct {
	URL  string
	Tags []string
}

func Path() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userHomeDir, ".config", "rss-tui", "feeds"), nil
}

func Load(path string) ([]FeedConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var feeds []FeedConfig

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		lineParts := strings.Fields(line)
		if len(lineParts) == 0 {
			continue
		}

		feed := FeedConfig{
			URL:  lineParts[0],
			Tags: lineParts[1:],
		}

		feeds = append(feeds, feed)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return feeds, nil
}
