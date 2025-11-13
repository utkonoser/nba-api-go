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

func TestGetTeamDashPtShots(t *testing.T) {
	mockResponse := `{
		"resource": "teamdashptshots",
		"parameters": {},
		"resultSets": [
{
				"name": "ClosestDefender10ftPlusShooting",
				"headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "G", "CLOSE_DEF_DIST_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ClosestDefenderShooting",
				"headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "G", "CLOSE_DEF_DIST_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DribbleShooting",
				"headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "G", "DRIBBLE_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "GeneralShooting",
				"headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "G", "SHOT_TYPE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotClockShooting",
				"headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "G", "SHOT_CLOCK_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TouchTimeShooting",
				"headers": ["TEAM_ID", "TEAM_NAME", "SORT_ORDER", "G", "TOUCH_TIME_RANGE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamdashptshots")
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

	params := TeamDashPtShotsParams{}

	response, err := c.GetTeamDashPtShots(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamdashptshots", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
