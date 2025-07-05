package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBase, err := url.Parse(rawBaseURL)
	if err != nil  {
		fmt.Println("Couldn't parse base url")
		return
	}
	parsedCurrent, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println("Couldn't parse current url")
		return
	}
	if parsedBase.Host != parsedCurrent.Host {
		return
	}
	normalizedCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println("Couldn't normalize current url")
		return
	}
	val, ok := pages[normalizedCurrent]
	if ok {
		pages[normalizedCurrent] = val + 1
	} else {
		pages[normalizedCurrent] = 1
		body, err := getHTML(rawCurrentURL)
		if err != nil {
			fmt.Printf("Couldn't get body for the current url: %v\n", err)
			return
		}
		urls, err := getURLsFromHTML(body, rawBaseURL)
		if err != nil {
			fmt.Println("Couldn't parse urls from the body")
			return
		}
		for _, v := range urls {
			crawlPage(rawBaseURL, v, pages)
		}
	}
}
