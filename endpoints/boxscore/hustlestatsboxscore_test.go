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

func TestGetHustleStatsBoxScore(t *testing.T) {
	mockResponse := `{
		"resource": "hustlestatsboxscore",
		"parameters": {},
		"resultSets": [
{
				"name": "HustleStatsAvailable",
				"headers": ["GAME_ID", "HUSTLE_STATUS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PlayerStats",
				"headers": ["GAME_ID", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_CITY", "PLAYER_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamStats",
				"headers": ["GAME_ID", "TEAM_ID", "TEAM_NAME", "TEAM_ABBREVIATION", "TEAM_CITY"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "hustlestatsboxscore")
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

	params := HustleStatsBoxScoreParams{}

	response, err := c.GetHustleStatsBoxScore(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "hustlestatsboxscore", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
