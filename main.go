package Flash

import (
	"fmt"
	"os"
)

func Flash() {
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

	// You can call your crawler logic here, e.g.
	// Flash.Crawl(baseURL)
}
