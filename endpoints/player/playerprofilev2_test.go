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

func TestGetPlayerProfileV2(t *testing.T) {
	mockResponse := `{
		"resource": "playerprofilev2",
		"parameters": {},
		"resultSets": [
{
				"name": "CareerHighs",
				"headers": ["PLAYER_ID", "GAME_DATE", "VS_TEAM_ID", "VS_TEAM_CITY", "VS_TEAM_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsAllStarSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "TEAM_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsCollegeSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "ORGANIZATION_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsPostSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "TEAM_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsPreseason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "TEAM_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CareerTotalsRegularSeason",
				"headers": ["PLAYER_ID", "LEAGUE_ID", "TEAM_ID", "GP", "GS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "NextGame",
				"headers": ["GAME_ID", "GAME_DATE", "GAME_TIME", "LOCATION", "PLAYER_TEAM_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "SeasonHighs",
				"headers": ["PLAYER_ID", "GAME_DATE", "VS_TEAM_ID", "VS_TEAM_CITY", "VS_TEAM_NAME"],
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
				"name": "SeasonTotalsPreseason",
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
		assert.Contains(t, r.URL.Path, "playerprofilev2")
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

	params := PlayerProfileV2Params{}

	response, err := c.GetPlayerProfileV2(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerprofilev2", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
