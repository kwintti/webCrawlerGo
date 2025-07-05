package main

import (
	"fmt"
	"os"
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
	url := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", url)
	fmt.Println(getHTML(url))
	// pages := map[string]int{}
	// crawlPage(url, url, pages)
	// for k, v := range pages {
	// 	fmt.Println(k, ": ", v)
	// }
}
