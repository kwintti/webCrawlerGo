package main

import (
	"net/url"
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
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	urls := []string{}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "" || a.Key[0] == '#' {
					continue
				}
				if a.Key == "href" {
					// if a.Val[0] == byte('/') {
					// 	urls = append(urls, rawBaseURL + a.Val)
					// 	break
					// }
					// urls = append(urls, a.Val)
					// break
					u, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					resolved := base.ResolveReference(u)
					if resolved.Host == base.Host {
						urls = append(urls, resolved.String())
					}
				}
			}
		}
	}

	return urls, nil
}
