package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	urlBase := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", urlBase)
	parsedUrlBase, err := url.Parse(urlBase)
	if err != nil {
		fmt.Println("Couldn't parse url")
		return
	}
	maxConcurrency := 1 
	config := config{
		baseUrl: parsedUrlBase,
		pages: map[string]int{},
		concurrencyControl:make(chan struct{},maxConcurrency),
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
	}
	config.crawlPage(parsedUrlBase.String())
	config.wg.Wait()
	for k, v := range config.pages {
		fmt.Println(k, ": ", v)
	}
}
