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

func TestGetScoreboard(t *testing.T) {
	mockResponse := `{
		"meta": {
			"version": 1,
			"request": "",
			"time": "2024-01-01T00:00:00Z",
			"code": 200
		},
		"scoreboard": {
			"gameDate": "2024-01-01",
			"leagueId": "00",
			"leagueName": "NBA",
			"games": [
				{
					"gameId": "0022400001",
					"gameCode": "20240101/LALLAC",
					"gameStatus": 1,
					"gameStatusText": "Final",
					"period": 4,
					"gameClock": "",
					"gameTimeUTC": "2024-01-01T01:00:00Z",
					"gameEt": "2024-01-01T20:00:00-05:00",
					"regulationPeriods": 4,
					"seriesGameNumber": "",
					"seriesText": "",
					"homeTeam": {
						"teamId": 1610612746,
						"teamName": "Clippers",
						"teamCity": "LA",
						"teamTricode": "LAC",
						"wins": 25,
						"losses": 15,
						"score": 110,
						"inBonus": null,
						"timeoutsRemaining": 2,
						"periods": [
							{
								"period": 1,
								"periodType": "REGULAR",
								"score": 28
							}
						]
					},
					"awayTeam": {
						"teamId": 1610612747,
						"teamName": "Lakers",
						"teamCity": "Los Angeles",
						"teamTricode": "LAL",
						"wins": 20,
						"losses": 20,
						"score": 105,
						"inBonus": null,
						"timeoutsRemaining": 1,
						"periods": [
							{
								"period": 1,
								"periodType": "REGULAR",
								"score": 25
							}
						]
					},
					"gameLeaders": {
						"homeLeaders": {
							"personId": 1629029,
							"name": "Player One",
							"jerseyNum": "2",
							"position": "G",
							"teamTricode": "LAC",
							"playerSlug": null,
							"points": 30,
							"rebounds": 8,
							"assists": 10
						},
						"awayLeaders": {
							"personId": 2544,
							"name": "Player Two",
							"jerseyNum": "23",
							"position": "F",
							"teamTricode": "LAL",
							"playerSlug": null,
							"points": 28,
							"rebounds": 7,
							"assists": 9
						}
					},
					"pbOdds": {
						"team": null,
						"odds": 0.0,
						"suspended": 0
					}
				}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "scoreboard/todaysScoreboard_00.json")
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

	scoreboard, err := c.GetScoreboard(context.Background())

	require.NoError(t, err)
	require.NotNil(t, scoreboard)
	assert.Equal(t, 200, scoreboard.Meta.Code)
	assert.Equal(t, "2024-01-01", scoreboard.Scoreboard.GameDate)
	assert.Equal(t, "NBA", scoreboard.Scoreboard.LeagueName)
	assert.Len(t, scoreboard.Scoreboard.Games, 1)

	game := scoreboard.Scoreboard.Games[0]
	assert.Equal(t, "0022400001", game.GameID)
	assert.Equal(t, "Final", game.GameStatusText)
	assert.Equal(t, 110, game.HomeTeam.Score)
	assert.Equal(t, 105, game.AwayTeam.Score)
	assert.Equal(t, "Clippers", game.HomeTeam.TeamName)
	assert.Equal(t, "Lakers", game.AwayTeam.TeamName)
}

func TestGetScoreboard_InvalidJSON(t *testing.T) {
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

	_, err := c.GetScoreboard(context.Background())

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid JSON")
}

func TestGetScoreboard_NetworkError(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	httpClient := client.NewHTTPClient("http://invalid-url-that-does-not-exist.local/%s", DefaultHeaders(), logger)
	c := &Client{
		httpClient: httpClient,
		logger:     logger,
	}

	_, err := c.GetScoreboard(context.Background())

	assert.Error(t, err)
}

