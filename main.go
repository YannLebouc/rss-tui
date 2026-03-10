package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"time"
)

func main(){
    log.SetPrefix("file reader: ")
    log.SetFlags(0)

    type rssURL struct {
    	url string
    	tag string
    }

    type channelImage struct {
    	url string
    	title string
    	link string
    	width int
    	height int
    	description string
    }

    type rssItem struct {
    	author string
    	categories []string
    	comments string
    	guid string
    	publicationDate time.Time
    	source string
    }

    type rssChannel struct {
		title string
		description string
		link string
		language string
		copyright string
		publicationDate time.Time
		lastBuildDate time.Time
		categories []string
		image channelImage
	}
	rssUrls := []rssURL {}

	file, err := os.Open("/home/ylebouc/dotfiles/common/rss-tui/feeds")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rssUrl := rssURL{
			url: scanner.Text(),
			tag: "",
		}
		rssUrls = append(rssUrls, rssUrl)
	}

	for _, value := range(rssUrls) {
		fmt.Println(value)
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
