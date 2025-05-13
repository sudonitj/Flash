package crawler

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Crawl starts the web crawling process from a base URL
func Crawl(baseURL string) {
	// Keep track of visited URLs to avoid revisiting
	visited := make(map[string]bool)
	
	// Start with the base URL
	pagesToVisit := []string{baseURL}
	
	// Simple crawling implementation
	for len(pagesToVisit) > 0 {
		// Get the next URL to visit
		url := pagesToVisit[0]
		pagesToVisit = pagesToVisit[1:] // Remove the URL from the queue
		
		// Skip if already visited
		normalized, err := NormalizeURL(url)
		if err != nil {
			fmt.Printf("Error normalizing URL %s: %v\n", url, err)
			continue
		}
		
		if visited[normalized] {
			continue
		}
		
		// Mark as visited
		visited[normalized] = true
		
		fmt.Printf("Crawling: %s\n", url)
		
		// Fetch the page
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			continue
		}
		
		// Check if the response is HTML
		contentType := resp.Header.Get("Content-Type")
		if !isHTML(contentType) {
			fmt.Printf("Skipping non-HTML content: %s (Content-Type: %s)\n", url, contentType)
			resp.Body.Close()
			continue
		}
		
		// Read the body
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Error reading body from %s: %v\n", url, err)
			continue
		}
		
		// Extract links from the page
		links, err := GetURLsFromHTML(string(body), url)
		if err != nil {
			fmt.Printf("Error extracting links from %s: %v\n", url, err)
			continue
		}
		
		// Add new links to the queue
		for _, link := range links {
			if !visited[link] {
				pagesToVisit = append(pagesToVisit, link)
			}
		}
		
		// Be nice to the server
		time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Printf("Crawling complete. Visited %d pages.\n", len(visited))
}

// isHTML checks if the content type indicates HTML
func isHTML(contentType string) bool {
	return contentType == "text/html" || contentType == "application/xhtml+xml" || 
		   contentType == "text/html; charset=utf-8" || contentType == "text/html; charset=UTF-8"
}