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

func TestGetPlayerDashboardByLastNGames(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashboardbylastngames",
		"parameters": {},
		"resultSets": [
{
				"name": "GameNumberPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last10PlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last15PlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last20PlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last5PlayerDashboard",
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
		assert.Contains(t, r.URL.Path, "playerdashboardbylastngames")
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

	params := PlayerDashboardByLastNGamesParams{}

	response, err := c.GetPlayerDashboardByLastNGames(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashboardbylastngames", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
