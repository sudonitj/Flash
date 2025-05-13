package crawler

import (
	"net/url"
	"strings"
)

// NormalizeURL normalizes a raw URL to a consistent format.
func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	hostname := parsedURL.Hostname()
	path := strings.TrimSuffix(parsedURL.EscapedPath(), "/")

	return hostname + path, nil
}