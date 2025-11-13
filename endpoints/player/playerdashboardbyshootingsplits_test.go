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

func TestGetPlayerDashboardByShootingSplits(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashboardbyshootingsplits",
		"parameters": {},
		"resultSets": [
{
				"name": "AssistedBy",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "FGM", "FGA"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "AssitedShotPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "OverallPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Shot5FTPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Shot8FTPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotAreaPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotTypePlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotTypeSummaryPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerdashboardbyshootingsplits")
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

	params := PlayerDashboardByShootingSplitsParams{}

	response, err := c.GetPlayerDashboardByShootingSplits(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashboardbyshootingsplits", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
