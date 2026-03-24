package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

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
