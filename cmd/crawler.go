// concCrawler provides a recursive web crawler as a binary/executable
// TODO: add documentation in doc.
// TODO: replace os.Args with flags.Parse
// TODO: add channel communication/concurrancy to crawler
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/DanielSchuette/webcrawler"
)

func main() {
	defer fmt.Printf("done\n%s\n",
		"------------------------------------------------------------")
	usagemsg := "\n'crawler' is a recursive web crawler.\n\nUsage:\n--help: Show this help message\n'crawler' takes a valid url (default: https://danielschuette.github.io)\nas its first input and a depth (default: 1) as its second input.\n\nWithout parameters, 'crawler' is executed with default values.\n\nAll urls that 'crawler' visits are saved in a hash map. Currently, this map cannot be retrieved.\nFuture versions might support saving all values to a specified file path.\n\nSee the GitHub repository at https://github.com/DanielSchuette for more information and updates.\n"
	var url string
	var depth int
	var CONVERR error
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Println(usagemsg)
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		url = os.Args[1]
	} else if len(os.Args) == 3 {
		depth, CONVERR = strconv.Atoi(os.Args[2])
		if CONVERR != nil {
			log.Fatalf("cannot convert flag to int: %s\n", CONVERR)
		}
		url = os.Args[1]
	} else {
		url = "http://danielschuette.github.io/"
		depth = 1
	}
	fmt.Printf("starting to crawl at %s...\n", url)
	for i := 0; i < 60; i++ {
		fmt.Printf("%s", "-")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("")
	// start shallow crawling
	shallowres, err := webcrawler.LinCrawl(url)
	if err != nil {
		fmt.Printf("error while crawling %s: %s\n", url, err)
	}
	if depth > 1 {
		// crawl multiple urls and write them to a map
		urlmap := make(map[string]string)
		urlmap = webcrawler.RecCrawl(shallowres, 1, depth, urlmap)
		fmt.Printf("\n\nmap after %d iterations (key='depth--url_number': value='url'):\n%+v\n\n\n", depth, urlmap)
	}
}
