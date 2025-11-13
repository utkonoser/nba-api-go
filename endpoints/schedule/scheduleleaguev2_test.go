package schedule

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

func TestGetScheduleLeagueV2(t *testing.T) {
	mockResponse := `{
		"resource": "scheduleleaguev2",
		"parameters": {},
		"resultSets": [
{
				"name": "SeasonGames",
				"headers": ["leagueId", "seasonYear", "gameDate", "gameId", "gameCode"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonWeeks",
				"headers": ["leagueId", "seasonYear", "weekNumber", "weekName", "startDate"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "scheduleleaguev2")
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

	params := ScheduleLeagueV2Params{}

	response, err := c.GetScheduleLeagueV2(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "scheduleleaguev2", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
