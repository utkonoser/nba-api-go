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

func TestGetTeamInfoCommon(t *testing.T) {
	mockResponse := `{
		"resource": "teaminfocommon",
		"parameters": {},
		"resultSets": [
{
				"name": "AvailableSeasons",
				"headers": ["SEASON_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamInfoCommon",
				"headers": ["TEAM_ID", "SEASON_YEAR", "TEAM_CITY", "TEAM_NAME", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamSeasonRanks",
				"headers": ["LEAGUE_ID", "SEASON_ID", "TEAM_ID", "PTS_RANK", "PTS_PG"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teaminfocommon")
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

	params := TeamInfoCommonParams{}

	response, err := c.GetTeamInfoCommon(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teaminfocommon", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
