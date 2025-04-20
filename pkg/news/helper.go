package news

import (
	"net/http"
	"time"
)

// getHttpClient returns a configured HTTP client with standard settings
func getHttpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

// getStandardHeaders returns a map of standard HTTP headers
func getStandardHeaders() map[string]string {
	return map[string]string{
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.5",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
	}
}
