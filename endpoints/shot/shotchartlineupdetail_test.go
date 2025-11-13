package shot

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

func TestGetShotChartLineupDetail(t *testing.T) {
	mockResponse := `{
		"resource": "shotchartlineupdetail",
		"parameters": {},
		"resultSets": [
{
				"name": "ShotChartLineupDetail",
				"headers": ["GRID_TYPE", "GAME_ID", "GAME_EVENT_ID", "GROUP_ID", "GROUP_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotChartLineupLeagueAverage",
				"headers": ["GRID_TYPE", "SHOT_ZONE_BASIC", "SHOT_ZONE_AREA", "SHOT_ZONE_RANGE", "FGA"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "shotchartlineupdetail")
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

	params := ShotChartLineupDetailParams{}

	response, err := c.GetShotChartLineupDetail(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "shotchartlineupdetail", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
