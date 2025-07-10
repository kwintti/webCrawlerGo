package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
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
	maxConcurrency, err := strconv.Atoi(os.Args[2]) 
	if err != nil {
		fmt.Println("Couldn't convert string to int")
	}
	maxPages, err := strconv.Atoi(os.Args[3]) 
	if err != nil {
		fmt.Println("Couldn't convert string to int")
	}
	config := config{
		baseUrl: parsedUrlBase,
		pages: map[string]int{},
		concurrencyControl:make(chan struct{},maxConcurrency),
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
		maxPages: maxPages,
	}
	config.wg.Add(1)
	config.crawlPage(parsedUrlBase.String())
	config.wg.Wait()
	printReport(config.pages, *config.baseUrl)

}
