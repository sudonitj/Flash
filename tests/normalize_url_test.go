package tests

import (
	"testing"

	"github.com/sudonitj/Flash"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme and trailing slash",
			inputURL: "https://blog.ekas.sudo/path/",
			expected: "blog.ekas.sudo/path",
		},
		{
			name:     "http scheme with trailing slash",
			inputURL: "http://blog.ekas.sudo/path/",
			expected: "blog.ekas.sudo/path",
		},
		{
			name:     "https scheme without trailing slash",
			inputURL: "https://blog.ekas.sudo/path",
			expected: "blog.ekas.sudo/path",
		},
		{
			name:     "http scheme without trailing slash",
			inputURL: "http://blog.ekas.sudo/path",
			expected: "blog.ekas.sudo/path",
		},
		{
			name:     "URL with no path",
			inputURL: "https://blog.ekas.sudo",
			expected: "blog.ekas.sudo",
		},
		{
			name:     "URL with path and query",
			inputURL: "https://blog.ekas.sudo/path?query=123",
			expected: "blog.ekas.sudo/path",
		},
		{
			name:     "URL with port number",
			inputURL: "https://blog.ekas.sudo:443/path",
			expected: "blog.ekas.sudo/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Flash.NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
