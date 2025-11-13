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

func TestGetPlayerDashboardByGameSplits(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashboardbygamesplits",
		"parameters": {},
		"resultSets": [
{
				"name": "ByActualMarginPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ByHalfPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ByPeriodPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ByScoreMarginPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "OverallPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerdashboardbygamesplits")
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

	params := PlayerDashboardByGameSplitsParams{}

	response, err := c.GetPlayerDashboardByGameSplits(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashboardbygamesplits", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
