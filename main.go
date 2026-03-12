package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

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

	// find the best way to get the path of the feeds file on any system, using os package to get the home directory for example
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
		feed := Feed{
			url:  scanner.Text(),
			tags: []string{},
		}
		feeds = append(feeds, feed)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Isolating an rss feed for easier implementation

	feed := feeds[0]
	fmt.Println(feed.url)
}

