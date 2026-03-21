package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ConfigLine struct {
	url  string
	tags []string
}

type ConfigFile struct {
	path  string
	lines []string
	line  ConfigLine
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

type Feed struct {
	Url     string   `xml:"-"`
	Tags    []string `xml:"-"`
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

func (c *ConfigFile) Path() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	c.path = filepath.Join(userHomeDir, ".config", "rss-tui", "feeds")
}

func (c *ConfigFile) Lines(path string) {
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

	c.lines = lines
}

func (c *ConfigFile) Parse(lines []string) {
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
		c.line.url = url
		c.line.tags = tags
	}
}

func FetchFeedsFromURL(feedUrl string) []byte {
	response, err := http.Get(feedUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode >= 300 {
		log.Fatalf("Response failed with status code %d and body %s\n", response.StatusCode, body)
	}

	return body
}

func InitializeFeeds(parsedConfigLines map[string][]string) (feeds []Feed) {
	for url, tags := range parsedConfigLines {
		feed := Feed{
			Url:  url,
			Tags: tags,
		}
		feeds = append(feeds, feed)
	}
	return feeds
}

func DecodeXML() {

}

func main() {
	configFile := ConfigFile{}
	configFile.Path()
	configFile.Lines(configFile.path)
	configFile.Parse(configFile.lines)
	feedsFromConfig := InitializeFeeds(configFile.Lines())

	fmt.Println(feedsFromConfig)
}
