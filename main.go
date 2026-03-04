package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main(){
	// fmt.Println("Hello World.")
    log.SetPrefix("file reader: ")
    log.SetFlags(0)

    type rssURL struct {
    	url string
    	tag string
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
