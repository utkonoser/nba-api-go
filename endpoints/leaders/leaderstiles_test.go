package leaders

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

func TestGetLeadersTiles(t *testing.T) {
	mockResponse := `{
		"resource": "leaderstiles",
		"parameters": {},
		"resultSets": [
{
				"name": "AllTimeSeasonHigh",
				"headers": ["TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "SEASON_YEAR", "PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "LastSeasonHigh",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "LeadersTiles",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "LowSeasonHigh",
				"headers": ["TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "SEASON_YEAR", "PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "leaderstiles")
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

	params := LeadersTilesParams{}

	response, err := c.GetLeadersTiles(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "leaderstiles", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
