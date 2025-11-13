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

func TestGetTeamPlayerDashboard(t *testing.T) {
	mockResponse := `{
		"resource": "teamplayerdashboard",
		"parameters": {},
		"resultSets": [
{
				"name": "PlayersSeasonTotals",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamOverall",
				"headers": ["GROUP_SET", "TEAM_ID", "TEAM_NAME", "GROUP_VALUE", "GP"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamplayerdashboard")
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

	params := TeamPlayerDashboardParams{}

	response, err := c.GetTeamPlayerDashboard(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamplayerdashboard", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
