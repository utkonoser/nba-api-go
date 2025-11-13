package live

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

func TestGetOdds(t *testing.T) {
	mockResponse := `{
		"meta": {
			"version": 1,
			"code": 200,
			"request": "",
			"time": "2024-01-01T00:00:00Z"
		},
		"game": {
			"gameId": "0022400001",
			"homeTeam": {
				"teamId": 1610612746,
				"teamName": "Clippers",
				"teamCity": "LA",
				"teamTricode": "LAC"
			},
			"awayTeam": {
				"teamId": 1610612747,
				"teamName": "Lakers",
				"teamCity": "Los Angeles",
				"teamTricode": "LAL"
			},
			"gameOdds": [
				{
					"provider": "DraftKings",
					"homeTeamOdds": {
						"moneyline": -150,
						"spread": -3.5,
						"spreadOdds": -110
					},
					"awayTeamOdds": {
						"moneyline": 130,
						"spread": 3.5,
						"spreadOdds": -110
					},
					"overUnder": {
						"total": 225.5,
						"overOdds": -110,
						"underOdds": -110
					},
					"suspended": 0
				},
				{
					"provider": "FanDuel",
					"homeTeamOdds": {
						"moneyline": -145,
						"spread": -3.0,
						"spreadOdds": -112
					},
					"awayTeamOdds": {
						"moneyline": 125,
						"spread": 3.0,
						"spreadOdds": -108
					},
					"overUnder": {
						"total": 226.0,
						"overOdds": -110,
						"underOdds": -110
					},
					"suspended": 0
				}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "odds/odds_0022400001.json")
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

	odds, err := c.GetOdds(context.Background(), "0022400001")

	require.NoError(t, err)
	require.NotNil(t, odds)
	assert.Equal(t, 200, odds.Meta.Code)
	assert.Equal(t, "0022400001", odds.Game.GameID)
	assert.Equal(t, "Clippers", odds.Game.HomeTeam.TeamName)
	assert.Equal(t, "Lakers", odds.Game.AwayTeam.TeamName)
	assert.Len(t, odds.Game.GameOdds, 2)

	draftKings := odds.Game.GameOdds[0]
	assert.Equal(t, "DraftKings", draftKings.Provider)
	assert.Equal(t, float64(-150), draftKings.HomeTeamOdds.Moneyline)
	assert.Equal(t, float64(130), draftKings.AwayTeamOdds.Moneyline)
	assert.Equal(t, float64(-3.5), draftKings.HomeTeamOdds.Spread)
	assert.Equal(t, float64(225.5), draftKings.OverUnder.Total)

	fanDuel := odds.Game.GameOdds[1]
	assert.Equal(t, "FanDuel", fanDuel.Provider)
	assert.Equal(t, float64(-145), fanDuel.HomeTeamOdds.Moneyline)
}

func TestGetOdds_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	httpClient := client.NewHTTPClient(server.URL+"/%s", DefaultHeaders(), logger)
	c := &Client{
		httpClient: httpClient,
		logger:     logger,
	}

	_, err := c.GetOdds(context.Background(), "0022400001")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid JSON")
}

