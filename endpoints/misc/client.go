// Package misc provides access to NBA Stats API misc endpoints.
package misc

import (
	"log/slog"

	"github.com/utkonoser/nba-api-go/client"
)

const (
	baseURL = "https://stats.nba.com/stats/%s"
)

// DefaultHeaders returns the default headers for NBA Stats API requests.
func DefaultHeaders() map[string]string {
	return map[string]string{
		"Host":            "stats.nba.com",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",
		"Accept":          "application/json, text/plain, */*",
		"Accept-Language": "en-US,en;q=0.5",
		"Accept-Encoding": "gzip, deflate, br",
		"Connection":      "keep-alive",
		"Referer":         "https://stats.nba.com/",
		"Pragma":          "no-cache",
		"Cache-Control":   "no-cache",
		"Sec-Ch-Ua":       `"Chromium";v="140", "Google Chrome";v="140", "Not;A=Brand";v="24"`,
		"Sec-Ch-Ua-Mobile": "?0",
		"Sec-Fetch-Dest":  "empty",
	}
}

// Client provides access to NBA Stats API misc endpoints.
type Client struct {
	httpClient *client.HTTPClient
	logger     *slog.Logger
}

// NewClient creates a new NBA Stats API misc client.
func NewClient(logger *slog.Logger) *Client {
	if logger == nil {
		logger = slog.Default()
	}

	httpClient := client.NewHTTPClient(baseURL, DefaultHeaders(), logger)

	return &Client{
		httpClient: httpClient,
		logger:     logger,
	}
}

// NewClientWithHeaders creates a new NBA Stats API misc client with custom headers.
func NewClientWithHeaders(headers map[string]string, logger *slog.Logger) *Client {
	if logger == nil {
		logger = slog.Default()
	}

	httpClient := client.NewHTTPClient(baseURL, headers, logger)

	return &Client{
		httpClient: httpClient,
		logger:     logger,
	}
}
