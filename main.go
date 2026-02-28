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

	file, err := os.Open("/home/ylebouc/dotfiles/rss-tui/feeds")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
