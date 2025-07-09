package main

import (
	"cmp"
	"fmt"
	"net/url"
	"slices"
	"strings"
)

type Urls struct {
	URL string
	Num int
	FirstLetter rune
}

func printReport(pages map[string]int, baseURL url.URL) {
	fmt.Println("=============================")
	fmt.Println("  REPORT for", baseURL.Scheme + "://" + baseURL.Host)
	fmt.Println("=============================")
	sorted := sortPages(pages, baseURL)

	for i := range sorted {
		fmt.Printf("Found %v internal links to %v\n", sorted[i].Num, sorted[i].URL)
	}

}

func sortPages(pages map[string]int, baseURL url.URL) []Urls {
	urls := make([]Urls, 0)
	for k, v := range pages {
		urls = append(urls, Urls{URL: k, Num: v, })
	}
	slices.SortFunc(urls, func(a, b Urls) int {
		if n := cmp.Compare(a.Num, b.Num); n != 0 {
			return n
		}
		return strings.Compare(a.URL, b.URL)
	}) 
	return urls 
}
