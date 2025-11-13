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

func TestGetPlayerDashboardByTeamPerformance(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashboardbyteamperformance",
		"parameters": {},
		"resultSets": [
{
				"name": "OverallPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PointsScoredPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE_ORDER", "GROUP_VALUE", "GROUP_VALUE_2", "GP"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PontsAgainstPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE_ORDER", "GROUP_VALUE", "GROUP_VALUE_2", "GP"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ScoreDifferentialPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE_ORDER", "GROUP_VALUE", "GROUP_VALUE_2", "GP"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerdashboardbyteamperformance")
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

	params := PlayerDashboardByTeamPerformanceParams{}

	response, err := c.GetPlayerDashboardByTeamPerformance(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashboardbyteamperformance", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
