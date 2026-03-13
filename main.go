package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ParseConfigurationLine(line string) (url string, tags []string) {

	if line == "" {
		return "", nil
	}

	if trimmedLine := strings.TrimSpace(line[0:1]); trimmedLine == "#" {
		return "", nil
	}

	stringParts := strings.Split(line, " ")
	if len(stringParts) > 0 {
		url = stringParts[0]
		if len(stringParts) > 1 {
			tags = stringParts[1:]
		}
	}

	return url, tags
}

func main() {
	log.SetPrefix("file reader: ")
	log.SetFlags(0)

	type ChannelImage struct {
		url         string
		title       string
		link        string
		width       int
		height      int
		description string
	}

	type Item struct {
		author          string
		categories      []string
		comments        string
		guid            string
		publicationDate time.Time
		source          string
	}

	type Channel struct {
		title           string
		description     string
		link            string
		language        string
		copyright       string
		publicationDate time.Time
		lastBuildDate   time.Time
		categories      []string
		image           ChannelImage
	}

	type Feed struct {
		url     string
		tags    []string
		channel Channel
		items   []Item
	}

	feeds := []Feed{}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(userHomeDir, ".config", "rss-tui", "feeds")

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url, tags := ParseConfigurationLine(scanner.Text())
		if url != "" && len(tags) > 0 {
			feed := Feed{
				url:  url,
				tags: tags,
			}
			feeds = append(feeds, feed)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.url)
		fmt.Println(feed.tags)
	}

	// Isolating an rss feed for easier implementation
}
