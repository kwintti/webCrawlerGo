package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	config := config{
		baseUrl: parsedUrlBase,
		pages: map[string]int{},
		concurrencyControl:make(chan struct{},maxConcurrency),
		mu: &sync.Mutex{},
		wg: &sync.WaitGroup{},
		maxPages: maxPages,
		ctx: ctx,
		ctxC: cancel,
	}
	config.wg.Add(1)
	fmt.Println("url base: ", parsedUrlBase.Path)
	config.crawlPage(parsedUrlBase.String())
	config.wg.Wait()
	printReport(config.pages, *config.baseUrl)

}
