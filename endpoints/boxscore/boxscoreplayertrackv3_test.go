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

func TestGetBoxScorePlayerTrackV3(t *testing.T) {
	mockResponse := `{
		"resource": "boxscoreplayertrackv3",
		"parameters": {},
		"resultSets": [
{
				"name": "PlayerStats",
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
		assert.Contains(t, r.URL.Path, "boxscoreplayertrackv3")
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

	params := BoxScorePlayerTrackV3Params{}

	response, err := c.GetBoxScorePlayerTrackV3(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "boxscoreplayertrackv3", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
