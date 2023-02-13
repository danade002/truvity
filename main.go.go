package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type URLResult struct {
	URL      string
	BodySize int
}

type URLResults []URLResult

func (r URLResults) Len() int {
	return len(r)
}

func (r URLResults) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r URLResults) Less(i, j int) bool {
	return r[i].BodySize < r[j].BodySize
}

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.amazon.com"}
	results := make(URLResults, 0, len(urls))

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error visiting %s: %s\n", url, err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response from %s: %s\n", url, err)
			continue
		}

		results = append(results, URLResult{
			URL:      url,
			BodySize: len(body),
		})
	}

	sort.Sort(results)

	for _, result := range results {
		fmt.Printf("%s: %d bytes\n", result.URL, result.BodySize)
	}
}
