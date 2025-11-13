package league

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/utkonoser/nba-api-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLeagueDashLineups(t *testing.T) {
	mockResponse := `{
		"resource": "leaguedashlineups",
		"parameters": {},
		"resultSets": [
{
				"name": "Lineups",
				"headers": ["GROUP_SET", "GROUP_ID", "GROUP_NAME", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "leaguedashlineups")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	httpClient := client.NewHTTPClient(server.URL+"/%s", DefaultHeaders(), logger)
	c := &Client{
		httpClient: httpClient,
		logger:     logger,
	}

	params := LeagueDashLineupsParams{}

	response, err := c.GetLeagueDashLineups(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "leaguedashlineups", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
