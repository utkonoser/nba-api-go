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

func TestGetTeamDashboardByGeneralSplits(t *testing.T) {
	mockResponse := `{
		"resource": "teamdashboardbygeneralsplits",
		"parameters": {},
		"resultSets": [
{
				"name": "DaysRestTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "TEAM_DAYS_REST_RANGE", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "LocationTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "TEAM_GAME_LOCATION", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "MonthTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "SEASON_MONTH_NAME", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "OverallTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "SEASON_YEAR", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PrePostAllStarTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "SEASON_SEGMENT", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "WinsLossesTeamDashboard",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GAME_RESULT", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamdashboardbygeneralsplits")
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

	params := TeamDashboardByGeneralSplitsParams{}

	response, err := c.GetTeamDashboardByGeneralSplits(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamdashboardbygeneralsplits", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
