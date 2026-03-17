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

	type Item struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		// Link        string `xml:"link"`
		// PubDate     time.Time `xml:"pubDate"`
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

	feeds := []Feed{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url, tags := ParseConfigurationLine(scanner.Text())
		if url != "" {
			feed := Feed{
				Url: url,
			}

			if len(tags) > 0 {
				feed.Tags = tags
			}

			feeds = append(feeds, feed)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	response, err := http.Get(feeds[0].Url)
	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and body %s\n", response.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	rssFeed := Feed{}

	xml.Unmarshal(body, &rssFeed)

	fmt.Println("CHANNEL TITLE : " + rssFeed.Channel.Title)
	fmt.Println("CHANNEL DESCRIPTION : " + rssFeed.Channel.Description)
	// fmt.Println(rssFeed.Channel.Items)
	for _, feed := range rssFeed.Channel.Items {
		fmt.Println("feed title : " + feed.Title)
		// fmt.Println("feed link : " + feed.Link)
		fmt.Println("feed description : " + feed.Description)
		// fmt.Println("feed pubDate : " + feed.PubDate.String())
	}
}
