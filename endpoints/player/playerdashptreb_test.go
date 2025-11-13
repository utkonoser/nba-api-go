package player

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

func TestGetPlayerDashPtReb(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashptreb",
		"parameters": {},
		"resultSets": [
{
				"name": "NumContestedRebounding",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SORT_ORDER", "G", "REB_NUM_CONTESTING_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "OverallRebounding",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "G", "OVERALL", "REB_FREQUENCY"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "RebDistanceRebounding",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SORT_ORDER", "G", "REB_DIST_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotDistanceRebounding",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SORT_ORDER", "G", "SHOT_DIST_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotTypeRebounding",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "SORT_ORDER", "G", "SHOT_TYPE_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerdashptreb")
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

	params := PlayerDashPtRebParams{}

	response, err := c.GetPlayerDashPtReb(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashptreb", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
