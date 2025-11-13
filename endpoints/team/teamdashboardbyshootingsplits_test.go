package team

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

func TestGetTeamDashboardByShootingSplits(t *testing.T) {
	mockResponse := `{
		"resource": "teamdashboardbyshootingsplits",
		"parameters": {},
		"resultSets": [
{
				"name": "AssistedBy",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "FGM", "FGA"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "AssitedShotTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "OverallTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Shot5FTTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Shot8FTTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotAreaTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotTypeTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "FGM", "FGA", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamdashboardbyshootingsplits")
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

	params := TeamDashboardByShootingSplitsParams{}

	response, err := c.GetTeamDashboardByShootingSplits(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamdashboardbyshootingsplits", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
