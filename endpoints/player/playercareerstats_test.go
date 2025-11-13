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

func TestGetPlayerCareerStats(t *testing.T) {
	mockResponse := `{
		"resource": "playercareerstats",
		"parameters": {},
		"resultSets": [
{
				"name": "CareerTotalsAllStarSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "Team_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsCollegeSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "ORGANIZATION_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsPostSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "Team_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsRegularSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "Team_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonRankingsPostSeason",
				"headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonRankingsRegularSeason",
				"headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonTotalsAllStarSeason",
				"headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonTotalsCollegeSeason",
				"headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "ORGANIZATION_ID", "SCHOOL_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonTotalsPostSeason",
				"headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonTotalsRegularSeason",
				"headers": ["PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playercareerstats")
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

	params := PlayerCareerStatsParams{}

	response, err := c.GetPlayerCareerStats(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playercareerstats", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
