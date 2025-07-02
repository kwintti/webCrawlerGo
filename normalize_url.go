package main

import (
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {
	parsed, err := url.Parse(u)	
	if err != nil {
		return "", err 
	}
	target := strings.TrimRight(parsed.Path, "/")
	normalized := parsed.Host + target

	return normalized, nil
}
