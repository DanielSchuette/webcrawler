package conccrawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

// Crawl takes a url and requests the respective http. A loop and the utility function findDelim() is used to find an <a ...> tag.
func Crawl(u string) []string {
	fmt.Printf("crawling %s\n", u)
	results := make([]string, 1)
	resp, err := http.Get(u)
	if err != nil {
		fmt.Printf("error during GET request: %s\n", err)
		return nil
	}
	// fmt.Printf("GET request result: %v\n", resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error while reading body: %s\n", err)
		return nil
	}
	// fmt.Printf("response BODY: %v\n", string(body[0:50]))
	for i := 0; i < len(body); i++ {
		if string(body[i:(i+6)]) == "a href" {
			l := string(body[(i + 8):(findDelim(body, i) + 1)])
			results = append(results, l)
		}
	}
	fmt.Printf("results slice: %+v\n", results)
	return results
}
