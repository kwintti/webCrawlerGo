package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages map[string]int
	baseUrl *url.URL
	mu *sync.Mutex
	concurrencyControl chan struct{}
	wg *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.wg.Add(1)
	cfg.concurrencyControl <- struct{}{}
	defer cfg.wg.Done()
	defer func(){<-cfg.concurrencyControl}()
	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Couldn't parse current url")
		return
	}
	if cfg.baseUrl.Host != parsedCurrent.Host {
		return
	}
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Couldn't normalize current url")
		return
	}
	cfg.mu.Lock()
	val, ok := cfg.pages[normalizedCurrent]
	cfg.mu.Unlock()
	if ok {
		cfg.mu.Lock()
		cfg.pages[normalizedCurrent] = val + 1
		cfg.mu.Unlock()
	} else {
		cfg.mu.Lock()
		cfg.pages[normalizedCurrent] = 1
		cfg.mu.Unlock()
		body, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Printf("Couldn't get body for the current url: %v\n", err)
			return
		}
		urls, err := getURLsFromHTML(body, cfg.baseUrl.String())
		if err != nil {
			fmt.Println("Couldn't parse urls from the body")
			return
		}
		for _, v := range urls {
			go cfg.crawlPage(v)
		}
	}
}
