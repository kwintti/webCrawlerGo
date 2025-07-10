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
	maxPages int
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func(){<-cfg.concurrencyControl}()
	defer cfg.wg.Done()

	cfg.mu.Lock()
	lenPages := len(cfg.pages)
	if lenPages >= cfg.maxPages {
	    cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Couldn't parse current url")
		return
	}

	if cfg.baseUrl.Host != parsedCurrent.Host {
		return
	}
	fmt.Println("Right now crawling page: ", rawCurrentURL)
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Couldn't normalize current url")
		return
	}

	cfg.mu.Lock()
	val, ok := cfg.pages[normalizedCurrent]
	if ok {
		cfg.pages[normalizedCurrent] = val + 1
		cfg.mu.Unlock()
	} else {
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
			cfg.concurrencyControl <- struct{}{}
			normalizedChild, err := normalizeURL(v)
			if err != nil {
				<-cfg.concurrencyControl 
				continue
			}
			cfg.mu.Lock()
			if _, ok := cfg.pages[normalizedChild]; ok {
				cfg.pages[normalizedChild]++
				cfg.mu.Unlock()
				<-cfg.concurrencyControl 
				continue 
			}
			if len(cfg.pages) >= cfg.maxPages {
				cfg.mu.Unlock()
				<-cfg.concurrencyControl
				return
			}

			cfg.pages[normalizedChild] = 1
			cfg.mu.Unlock()
			cfg.wg.Add(1)
			go cfg.crawlPage(v)

		}
	}
}
