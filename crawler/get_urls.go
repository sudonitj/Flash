package crawler

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// GetURLsFromHTML parses HTML and extracts all anchor tag hrefs, converting them to absolute URLs.
func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := []string{} // Initialize empty slice for URLs

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href := attr.Val
					parsedURL, err := url.Parse(href)
					if err != nil {
						continue
					}
					finalURL := base.ResolveReference(parsedURL)
					urls = append(urls, finalURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)

	return urls, nil
}