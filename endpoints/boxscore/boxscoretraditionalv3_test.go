package boxscore

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

func TestGetBoxScoreTraditionalV3(t *testing.T) {
	mockResponse := `{
		"resource": "boxscoretraditionalv3",
		"parameters": {},
		"resultSets": [
{
				"name": "PlayerStats",
				"headers": ["gameId", "teamId", "teamCity", "teamName", "teamTricode"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamStarterBenchStats",
				"headers": ["gameId", "teamId", "teamCity", "teamName", "teamTricode"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamStats",
				"headers": ["gameId", "teamId", "teamCity", "teamName", "teamTricode"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "boxscoretraditionalv3")
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

	params := BoxScoreTraditionalV3Params{}

	response, err := c.GetBoxScoreTraditionalV3(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "boxscoretraditionalv3", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
