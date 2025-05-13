package main

import (
	"fmt"
	"os"
	
	"github.com/sudonitj/Flash/crawler"
)

func main() {
	args := os.Args[1:] // skip the program name

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	fmt.Println("starting crawl of:", baseURL)

	// Call crawler logic
	crawler.Crawl(baseURL)
}