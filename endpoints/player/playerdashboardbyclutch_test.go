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

func TestGetPlayerDashboardByClutch(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashboardbyclutch",
		"parameters": {},
		"resultSets": [
{
				"name": "Last10Sec3Point2PlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last10Sec3PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last1Min5PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last1MinPlusMinus5PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last30Sec3Point2PlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last30Sec3PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last3Min5PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last3MinPlusMinus5PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last5Min5PointPlayerDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Last5MinPlusMinus5PointPlayerDashboard",
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
		assert.Contains(t, r.URL.Path, "playerdashboardbyclutch")
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

	params := PlayerDashboardByClutchParams{}

	response, err := c.GetPlayerDashboardByClutch(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashboardbyclutch", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
