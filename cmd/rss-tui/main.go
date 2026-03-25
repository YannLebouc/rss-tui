package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"github.com/YannLebouc/rss-tui/internal/config"
)

func main() {
	configPath, err := config.Path()
	if err != nil {
		log.Fatal(err)
	}

	feeds, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

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
