package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
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

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type FeedConfig struct {
	URL  string
	Tags []string
}

type Feed struct {
	URL     string   `xml:"-"`
	Tags    []string `xml:"-"`
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type FetchedFeed struct {
	Config FeedConfig
	Feed   Feed
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

func InitializeFeeds(config *ConfigFile) []Feed {
	feeds := []Feed{}
	for _, line := range config.Lines {
		feed := Feed{
			URL:  line.URL,
			Tags: line.Tags,
		}
		feeds = append(feeds, feed)
	}
	return feeds
}

func main() {
	configFile := ConfigFile{
		Path: ConfigPath(),
	}
	configFile.Load()

	feeds := InitializeFeeds(&configFile)

	for i := range feeds {
		response, err := http.Get(feeds[i].URL)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode >= 300 {
			log.Fatalf("Response failed with status code %d and body %s\n", response.StatusCode)
		}

		decoder := xml.NewDecoder(response.Body)
		if err := decoder.Decode(&feeds[i]); err != nil {
			log.Fatalf("Failed to decode feed from URL %s: %v\n", feeds[i].URL, err)
		}
	}

	fmt.Println(feeds[0].Channel)
}
