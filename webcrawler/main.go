package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Crawler struct {
	seen map[string]error
	sync.RWMutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *Crawler) Crawl(url string, depth int, fetcher Fetcher) {
	done := make(chan bool)

	go func() {
		defer func() {
			done <- true
		}()

		if depth <= 0 {
			return
		}

		c.Lock()
		if _, ok := c.seen[url]; ok {
			c.Unlock()
			fmt.Printf("already seen: %s\n", url)
			return
		}

		c.seen[url] = nil
		c.Unlock()

		body, urls, err := fetcher.Fetch(url)
		fmt.Printf("found: %s %q\n", url, body)

		if err != nil {
			c.Lock()
			c.seen[url] = err
			c.Unlock()

			fmt.Println(err)
			return
		}

		for _, u := range urls {
			c.Crawl(u, depth-1, fetcher)
		}
	}()
	<-done
}

func NewCrawler() *Crawler {
	return &Crawler{seen: make(map[string]error)}
}

func main() {
	c := NewCrawler()
	c.Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",

			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
			"https://golang.org/pkg/sync/",
			"https://golang.org/pkg/atomic/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
