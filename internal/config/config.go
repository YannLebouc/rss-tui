package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ConfigLine struct {
	URL  string
	Tags []string
}

type ConfigFile struct {
	Path  string
	Lines []ConfigLine
}

type FeedConfig struct {
	URL  string
	Tags []string
}

func ConfigPath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(userHomeDir, ".config", "rss-tui", "feeds")
}

func ReadConfig(path string) []string {
	lines := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func ParseConfig(lines []string) []ConfigLine {
	parsedLines := []ConfigLine{}

	for _, line := range lines {
		var url string
		var tags []string
		stringParts := strings.Fields(line)
		if len(stringParts) > 0 {
			url = stringParts[0]
			if len(stringParts) > 1 {
				tags = stringParts[1:]
			}
		}

		parsedLine := ConfigLine{
			URL:  url,
			Tags: tags,
		}
		parsedLines = append(parsedLines, parsedLine)
	}

	return parsedLines
}

func (c *ConfigFile) Load() {
	rawLines := ReadConfig(c.Path)
	c.Lines = ParseConfig(rawLines)
}

func InitializeFeedsConfig(config *ConfigFile) []FeedConfig {
	feeds := []FeedConfig{}
	for _, line := range config.Lines {
		feed := FeedConfig{
			URL:  line.URL,
			Tags: line.Tags,
		}
		feeds = append(feeds, feed)
	}
	return feeds
}
