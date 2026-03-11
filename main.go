package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetPrefix("file reader: ")
	log.SetFlags(0)

	type FeedURL struct {
		url string
		tag string
	}

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
		rssUrl  FeedURL
		channel Channel
		items   []Item
	}

	rssUrls := []FeedURL{}

	file, err := os.Open("/home/ylebouc/dotfiles/common/rss-tui/feeds")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rssUrl := FeedURL{
			url: scanner.Text(),
			tag: "",
		}
		rssUrls = append(rssUrls, rssUrl)
	}

	// Isolating an rss feed for easier implementation

	rssUrl := rssUrls[0]
	fmt.Println(rssUrl.url)
	// for _, value := range rssUrls {
	// 	fmt.Println(value)
	// }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
