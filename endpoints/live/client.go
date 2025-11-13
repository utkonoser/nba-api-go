// Package live provides access to NBA Live Data API endpoints.
package live

import (
	"log/slog"

	"github.com/utkonoser/nba-api-go/client"
)

const (
	baseURL = "https://cdn.nba.com/static/json/liveData/%s"
)

// DefaultHeaders returns the default headers for NBA Live API requests.
func DefaultHeaders() map[string]string {
	return map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "en-US,en;q=0.9",
		"Cache-Control":   "max-age=0",
		"Connection":      "keep-alive",
		"Host":            "cdn.nba.com",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
	}
}

// Client provides access to NBA Live API endpoints.
type Client struct {
	httpClient *client.HTTPClient
	logger     *slog.Logger
}

// NewClient creates a new NBA Live API client.
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

// NewClientWithHeaders creates a new NBA Live API client with custom headers.
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

