package main

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	r := strings.NewReader(htmlBody)
	doc, err := html.Parse(r)	
	if err != nil {
		return nil, err
	}

	urls := []string{}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if a.Val[0] == byte('/') {
						urls = append(urls, rawBaseURL + a.Val)
						break
					}
					urls = append(urls, a.Val)
					break
				}
			}
		}
	}

	return urls, nil
}
