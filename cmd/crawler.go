// concCrawler provides a concurrent web crawler as a binary/executable
// TODO: add documentation in doc.go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/DanielSchuette/conccrawler"
)

func main() {
	defer fmt.Printf("done\n%s\n",
		"------------------------------------------------------------")
	var url string
	// TODO: replace os.Args with flags.Parse
	if len(os.Args) > 1 {
		url = os.Args[1]
	} else {
		url = "http://danielschuette.github.io/"
	}
	fmt.Printf("starting to crawl at %s...\n", url)
	for i := 0; i < 60; i++ {
		fmt.Printf("%s", "-")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("")
	// start crawling, TODO: crawl recursively and concurrently
	result := conccrawler.Crawl(url)
	for _, res := range result {
		_ = conccrawler.Crawl(res)
	}
}
