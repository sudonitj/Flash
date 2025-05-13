package crawler_test

import (
	"reflect"
	"testing"

	"github.com/sudonitj/Flash/crawler"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://ekas.sudo",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Ekas Link</span>
					</a>
					<a href="https://other.com/path/one">
						<span>External</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://ekas.sudo/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no links in body",
			inputURL: "https://ekas.sudo",
			inputBody: `
			<html><body><h1>No links here</h1></body></html>
			`,
			expected: []string{},
		},
		{
			name:     "multiple relative links",
			inputURL: "https://ekas.sudo",
			inputBody: `
			<html>
				<body>
					<a href="/alpha">A</a>
					<a href="/beta">B</a>
				</body>
			</html>
			`,
			expected: []string{"https://ekas.sudo/alpha", "https://ekas.sudo/beta"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := crawler.GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}