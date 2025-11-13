// Package client provides HTTP client functionality for NBA API.
package client

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
)

// HTTPClient provides HTTP functionality for making NBA API requests.
type HTTPClient struct {
	baseURL    string
	headers    map[string]string
	httpClient *http.Client
	logger     *slog.Logger
}

// Response represents an NBA API response.
type Response struct {
	raw        string
	statusCode int
	url        string
}

// NewHTTPClient creates a new HTTP client with the specified base URL and headers.
func NewHTTPClient(baseURL string, headers map[string]string, logger *slog.Logger) *HTTPClient {
	if logger == nil {
		logger = slog.Default()
	}

	return &HTTPClient{
		baseURL: baseURL,
		headers: headers,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		logger: logger,
	}
}

// SetTimeout sets the HTTP client timeout.
func (c *HTTPClient) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}

// SendRequest sends an HTTP GET request to the specified endpoint with parameters.
func (c *HTTPClient) SendRequest(ctx context.Context, endpoint string, params map[string]string) (*Response, error) {
	// Build URL with parameters
	fullURL := c.baseURL
	if endpoint != "" {
		fullURL = fmt.Sprintf(c.baseURL, endpoint)
	}

	// Sort parameters by key (NBA API sometimes requires sorted parameters)
	sortedParams := c.sortParameters(params)
	queryParams := c.buildQueryParams(sortedParams)

	if queryParams != "" {
		fullURL = fullURL + "?" + queryParams
	}

	c.logger.DebugContext(ctx, "Sending NBA API request",
		slog.String("url", fullURL),
		slog.Any("params", params))

	// Create request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to create request",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to send request",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body (handle gzip encoding)
	var reader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			c.logger.ErrorContext(ctx, "Failed to create gzip reader",
				slog.String("error", err.Error()))
			return nil, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to read response body",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	response := &Response{
		raw:        string(body),
		statusCode: resp.StatusCode,
		url:        fullURL,
	}

	c.logger.DebugContext(ctx, "Received NBA API response",
		slog.Int("status_code", resp.StatusCode),
		slog.Int("body_length", len(body)))

	return response, nil
}

// sortParameters sorts parameters by key.
func (c *HTTPClient) sortParameters(params map[string]string) []struct {
	key   string
	value string
} {
	var sorted []struct {
		key   string
		value string
	}

	for k, v := range params {
		sorted = append(sorted, struct {
			key   string
			value string
		}{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].key < sorted[j].key
	})

	return sorted
}

// buildQueryParams builds query parameters string.
func (c *HTTPClient) buildQueryParams(params []struct {
	key   string
	value string
}) string {
	values := url.Values{}
	for _, p := range params {
		if p.value != "" {
			values.Add(p.key, p.value)
		}
	}
	return values.Encode()
}

// GetRaw returns the raw response string.
func (r *Response) GetRaw() string {
	return r.raw
}

// GetStatusCode returns the HTTP status code.
func (r *Response) GetStatusCode() int {
	return r.statusCode
}

// GetURL returns the request URL.
func (r *Response) GetURL() string {
	return r.url
}

// GetJSON unmarshals the response into the provided interface.
func (r *Response) GetJSON(v interface{}) error {
	if err := json.Unmarshal([]byte(r.raw), v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return nil
}

// IsValidJSON checks if the response is valid JSON.
func (r *Response) IsValidJSON() bool {
	var js interface{}
	return json.Unmarshal([]byte(r.raw), &js) == nil
}

