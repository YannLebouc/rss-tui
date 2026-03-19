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
	"time"
)

func GetFeedsConfigFilePath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(userHomeDir, ".config", "rss-tui", "feeds")
	return filepath
}

func GetConfigFileLines(filepath string) (lines []string) {
	file, err := os.Open(filepath)
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

		if trimmedLine := strings.TrimSpace(line[0:1]); trimmedLine == "#" {
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ParseConfigurationLine(line string) (url string, tags []string) {
	stringParts := strings.Split(line, " ")
	if len(stringParts) > 0 {
		url = stringParts[0]
		if len(stringParts) > 1 {
			tags = stringParts[1:]
		}
	}
	return url, tags
}

func FetchFeedsFromURL(feedUrl string) []byte {
	response, err := http.Get(feedUrl)
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and body %s\n", response.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func DecodeXML() {

}

type Item struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	PubDate     time.Time `xml:"pubDate"`
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

func main() {

	filepath := GetFeedsConfigFilePath()
	feedUrls := GetConfigFileLines(filepath)
	feeds := []Feed{}
	for _, feed := range feedUrls {
		url, tags := ParseConfigurationLine(feed)
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

	rssFeed := Feed{}

	if err := xml.Unmarshal(body, &rssFeed); err != nil {
		log.Fatal(err)
	}

	fmt.Println("CHANNEL TITLE : " + rssFeed.Channel.Title)
	fmt.Println("CHANNEL DESCRIPTION : " + rssFeed.Channel.Description)
	fmt.Println(rssFeed.Channel.Items)
	// for _, feed := range rssFeed.Channel.Items {
	// 	fmt.Println("feed title : " + feed.Title)
	// 	fmt.Println("feed link : " + feed.Link)
	// 	fmt.Println("feed description : " + feed.Description)
	// 	fmt.Println("feed pubDate : " + feed.PubDate.String())
	// }
}
