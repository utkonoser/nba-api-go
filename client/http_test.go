package client

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPClient(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	headers := map[string]string{
		"User-Agent": "test-agent",
	}

	client := NewHTTPClient("https://api.example.com", headers, logger)

	assert.NotNil(t, client)
	assert.Equal(t, "https://api.example.com", client.baseURL)
	assert.Equal(t, headers, client.headers)
	assert.NotNil(t, client.httpClient)
	assert.NotNil(t, client.logger)
}

func TestNewHTTPClient_NilLogger(t *testing.T) {
	client := NewHTTPClient("https://api.example.com", nil, nil)

	assert.NotNil(t, client)
	assert.NotNil(t, client.logger)
}

func TestHTTPClient_SetTimeout(t *testing.T) {
	client := NewHTTPClient("https://api.example.com", nil, nil)
	timeout := 10 * time.Second

	client.SetTimeout(timeout)

	assert.Equal(t, timeout, client.httpClient.Timeout)
}

func TestHTTPClient_SendRequest(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/test", r.URL.Path)
		assert.Equal(t, "value1", r.URL.Query().Get("param1"))
		assert.Equal(t, "value2", r.URL.Query().Get("param2"))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := NewHTTPClient(server.URL+"/test", nil, nil)
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	resp, err := client.SendRequest(context.Background(), "", params)

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.GetStatusCode())
	assert.True(t, resp.IsValidJSON())
}

func TestHTTPClient_SendRequest_WithEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/test-endpoint", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data": "test"}`))
	}))
	defer server.Close()

	client := NewHTTPClient(server.URL+"/api/%s", nil, nil)

	resp, err := client.SendRequest(context.Background(), "test-endpoint", nil)

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.GetStatusCode())
}

func TestHTTPClient_SendRequest_WithHeaders(t *testing.T) {
	expectedHeaders := map[string]string{
		"User-Agent":   "test-agent",
		"Content-Type": "application/json",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range expectedHeaders {
			assert.Equal(t, value, r.Header.Get(key))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHTTPClient(server.URL, expectedHeaders, nil)

	_, err := client.SendRequest(context.Background(), "", nil)

	require.NoError(t, err)
}

func TestResponse_GetJSON(t *testing.T) {
	jsonData := `{"name": "test", "value": 123}`
	resp := &Response{
		raw:        jsonData,
		statusCode: 200,
		url:        "http://example.com",
	}

	var result map[string]interface{}
	err := resp.GetJSON(&result)

	require.NoError(t, err)
	assert.Equal(t, "test", result["name"])
	assert.Equal(t, float64(123), result["value"])
}

func TestResponse_GetJSON_InvalidJSON(t *testing.T) {
	resp := &Response{
		raw: "invalid json",
	}

	var result map[string]interface{}
	err := resp.GetJSON(&result)

	assert.Error(t, err)
}

func TestResponse_IsValidJSON(t *testing.T) {
	tests := []struct {
		name     string
		raw      string
		expected bool
	}{
		{
			name:     "valid JSON object",
			raw:      `{"key": "value"}`,
			expected: true,
		},
		{
			name:     "valid JSON array",
			raw:      `[1, 2, 3]`,
			expected: true,
		},
		{
			name:     "invalid JSON",
			raw:      `{invalid}`,
			expected: false,
		},
		{
			name:     "empty string",
			raw:      ``,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &Response{raw: tt.raw}
			assert.Equal(t, tt.expected, resp.IsValidJSON())
		})
	}
}

func TestHTTPClient_sortParameters(t *testing.T) {
	client := NewHTTPClient("", nil, nil)
	params := map[string]string{
		"zebra":  "z",
		"alpha":  "a",
		"mike":   "m",
		"bravo":  "b",
		"delta":  "d",
		"charlie": "c",
	}

	sorted := client.sortParameters(params)

	assert.Len(t, sorted, 6)
	assert.Equal(t, "alpha", sorted[0].key)
	assert.Equal(t, "bravo", sorted[1].key)
	assert.Equal(t, "charlie", sorted[2].key)
	assert.Equal(t, "delta", sorted[3].key)
	assert.Equal(t, "mike", sorted[4].key)
	assert.Equal(t, "zebra", sorted[5].key)
}

