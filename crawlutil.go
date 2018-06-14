package webcrawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// findDelim takes a slice of bytes t, a string d and a start value s that must be a valid index of t
// findDelim returns an int i that is the index of the next occurance of the delimiter after start value s (+8 bytes,
// otherwise findDelim will return the START and not the END index of an attribute value (i.e. ' " ')
func findDelim(t []byte, s int) int {
	ind := (s + 8)
	for i := ind; i < len(t); i++ {
		if string(t[i]) == "\"" {
			ind := i
			return (ind - 1) // don't want the " itself
		}
	}
	return (ind - 1) // don't want the " itself
}

// arrToMap writes an slice to an existing map
func sliceToMap(m map[string]string, arr []string, ind int) map[string]string {
	for idx, url := range arr {
		m[strconv.Itoa(ind)+"--"+strconv.Itoa(idx)] = url
	}
	return m
}

// LinCrawl takes a slice of urls and make a GET request via the respective protocol. A loop and the utility function findDelim() is used to find an <a ...> tag.
func LinCrawl(url string) ([]string, error) {
	results := make([]string, 0)
	// print the name of the currently crawled url
	fmt.Printf("crawling %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		e := fmt.Errorf("error during GET request: %s", err)
		return nil, e
	}
	// fmt.Printf("GET request result: %v\n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e := fmt.Errorf("error while reading body: %s", err)
		return nil, e
	}
	// fmt.Printf("response BODY: %v\n", string(body[0:50]))
	for i := 0; i < len(body); i++ {
		if string(body[i:(i+6)]) == "a href" {
			l := string(body[(i + 8):(findDelim(body, i) + 1)])
			results = append(results, l)
		}
	}
	fmt.Printf("results slice: %+v\n", results)
	return results, nil
}

// RecCrawl recursively crawls all links it detects to a certain specified depth while printing all the relevant information (uses the linear crawler as a helper!)
// to the screen. The map implementation is not ideal especially because maps can not safely be accessed via channels
// that means concurrent use is not yet supported. TODO: channel implementation of the crawler to enable concurrance
func RecCrawl(urls []string, current, depth int, visited map[string]string) map[string]string {
	if current == depth {
		return visited
	}
	current++ // increment the depth tracking value
	results := make([]string, 0)
	for idx, url := range urls {
		fmt.Printf("depth %d: crawling %d of %d\n", current, idx+1, len(urls))
		res, err := LinCrawl(url)
		if err != nil {
			fmt.Printf("error while crawling %s: %s\n", url, err)
			continue
		}
		for _, r := range res {
			results = append(results, r)
		}
	}
	visited = sliceToMap(visited, results, current)
	return RecCrawl(results, current, depth, visited)
}
