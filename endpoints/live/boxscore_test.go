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

func TestGetBoxScore(t *testing.T) {
	mockResponse := `{
		"meta": {
			"version": 1,
			"code": 200,
			"request": "",
			"time": "2024-01-01T00:00:00Z"
		},
		"game": {
			"gameId": "0022400001",
			"gameTimeLocal": "2024-01-01T19:30:00",
			"gameTimeUTC": "2024-01-02T00:30:00Z",
			"gameTimeHome": "2024-01-01T19:30:00-05:00",
			"gameTimeAway": "2024-01-01T19:30:00-05:00",
			"gameEt": "2024-01-01T19:30:00-05:00",
			"duration": 125,
			"gameCode": "20240101/LALLAC",
			"gameStatusText": "Final",
			"gameStatus": 3,
			"regulationPeriods": 4,
			"period": 4,
			"gameClock": "PT00M00.00S",
			"attendance": 19000,
			"sellout": "1",
			"arena": {
				"arenaId": 5,
				"arenaName": "Crypto.com Arena",
				"arenaCity": "Los Angeles",
				"arenaState": "CA",
				"arenaCountry": "US",
				"arenaTimezone": "America/Los_Angeles"
			},
			"officials": [
				{
					"personId": 201638,
					"name": "John Doe",
					"nameI": "J. Doe",
					"firstName": "John",
					"familyName": "Doe",
					"jerseyNum": "36",
					"assignment": "OFFICIAL1"
				}
			],
			"homeTeam": {
				"teamId": 1610612746,
				"teamName": "Clippers",
				"teamCity": "LA",
				"teamTricode": "LAC",
				"score": 110,
				"inBonus": "1",
				"timeoutsRemaining": 2,
				"periods": [
					{
						"period": 1,
						"periodType": "REGULAR",
						"score": 28
					}
				],
				"players": [
					{
						"status": "ACTIVE",
						"order": 1,
						"personId": 1629029,
						"jerseyNum": "2",
						"position": "G",
						"starter": "1",
						"oncourt": "0",
						"played": "1",
						"statistics": {
							"assists": 10,
							"blocks": 0,
							"blocksReceived": 0,
							"fieldGoalsAttempted": 20,
							"fieldGoalsMade": 12,
							"fieldGoalsPercentage": 0.6,
							"foulsOffensive": 0,
							"foulsDrawn": 4,
							"foulsPersonal": 2,
							"foulsTechnical": 0,
							"freeThrowsAttempted": 8,
							"freeThrowsMade": 6,
							"freeThrowsPercentage": 0.75,
							"minus": 50.0,
							"minutes": "PT35M00.00S",
							"minutesCalculated": "PT35M",
							"plus": 65.0,
							"plusMinusPoints": 15.0,
							"points": 30,
							"pointsFastBreak": 4,
							"pointsInThePaint": 10,
							"pointsSecondChance": 2,
							"reboundsDefensive": 5,
							"reboundsOffensive": 2,
							"reboundsTotal": 7,
							"steals": 3,
							"threePointersAttempted": 8,
							"threePointersMade": 0,
							"threePointersPercentage": 0.0,
							"turnovers": 2,
							"twoPointersAttempted": 12,
							"twoPointersMade": 12,
							"twoPointersPercentage": 1.0
						},
						"name": "Player One",
						"nameI": "P. One",
						"firstName": "Player",
						"familyName": "One"
					}
				],
				"statistics": {
					"assists": 25,
					"assistsTurnoverRatio": 2.5,
					"benchPoints": 40,
					"biggestLead": 15,
					"biggestLeadScore": "95-80",
					"biggestScoringRun": 10,
					"biggestScoringRunScore": "95-80",
					"blocks": 5,
					"blocksReceived": 3,
					"fastBreakPointsAttempted": 10,
					"fastBreakPointsMade": 7,
					"fastBreakPointsPercentage": 0.7,
					"fieldGoalsAttempted": 90,
					"fieldGoalsEffectiveAdjusted": 0.55,
					"fieldGoalsMade": 45,
					"fieldGoalsPercentage": 0.5,
					"foulsOffensive": 1,
					"foulsDrawn": 20,
					"foulsPersonal": 18,
					"foulsTeam": 18,
					"foulsTechnical": 0,
					"foulsTeamTechnical": 0,
					"freeThrowsAttempted": 25,
					"freeThrowsMade": 20,
					"freeThrowsPercentage": 0.8,
					"leadChanges": 5,
					"minutes": "PT240M00.00S",
					"minutesCalculated": "PT240M",
					"points": 110,
					"pointsAgainst": 105,
					"pointsFastBreak": 14,
					"pointsFromTurnovers": 18,
					"pointsInThePaint": 50,
					"pointsInThePaintAttempted": 40,
					"pointsInThePaintMade": 25,
					"pointsInThePaintPercentage": 0.625,
					"pointsSecondChance": 12,
					"reboundsDefensive": 35,
					"reboundsOffensive": 10,
					"reboundsPersonal": 45,
					"reboundsTeam": 5,
					"reboundsTeamDefensive": 2,
					"reboundsTeamOffensive": 3,
					"reboundsTotal": 50,
					"secondChancePointsAttempted": 10,
					"secondChancePointsMade": 6,
					"secondChancePointsPercentage": 0.6,
					"steals": 8,
					"threePointersAttempted": 35,
					"threePointersMade": 15,
					"threePointersPercentage": 0.428571,
					"timeLeading": "PT35M00.00S",
					"timesTied": 2,
					"trueShootingAttempts": 97.5,
					"trueShootingPercentage": 0.565,
					"turnovers": 10,
					"turnoversTeam": 0,
					"turnoversTotal": 10,
					"twoPointersAttempted": 55,
					"twoPointersMade": 30,
					"twoPointersPercentage": 0.545454
				}
			},
			"awayTeam": {
				"teamId": 1610612747,
				"teamName": "Lakers",
				"teamCity": "Los Angeles",
				"teamTricode": "LAL",
				"score": 105,
				"inBonus": "1",
				"timeoutsRemaining": 1,
				"periods": [
					{
						"period": 1,
						"periodType": "REGULAR",
						"score": 25
					}
				],
				"players": [],
				"statistics": {
					"assists": 20,
					"assistsTurnoverRatio": 1.67,
					"benchPoints": 35,
					"biggestLead": 5,
					"biggestLeadScore": "20-15",
					"biggestScoringRun": 8,
					"biggestScoringRunScore": "20-15",
					"blocks": 3,
					"blocksReceived": 5,
					"fastBreakPointsAttempted": 8,
					"fastBreakPointsMade": 5,
					"fastBreakPointsPercentage": 0.625,
					"fieldGoalsAttempted": 88,
					"fieldGoalsEffectiveAdjusted": 0.52,
					"fieldGoalsMade": 42,
					"fieldGoalsPercentage": 0.477,
					"foulsOffensive": 2,
					"foulsDrawn": 18,
					"foulsPersonal": 20,
					"foulsTeam": 20,
					"foulsTechnical": 1,
					"foulsTeamTechnical": 0,
					"freeThrowsAttempted": 22,
					"freeThrowsMade": 17,
					"freeThrowsPercentage": 0.773,
					"leadChanges": 5,
					"minutes": "PT240M00.00S",
					"minutesCalculated": "PT240M",
					"points": 105,
					"pointsAgainst": 110,
					"pointsFastBreak": 10,
					"pointsFromTurnovers": 12,
					"pointsInThePaint": 46,
					"pointsInThePaintAttempted": 38,
					"pointsInThePaintMade": 23,
					"pointsInThePaintPercentage": 0.605,
					"pointsSecondChance": 10,
					"reboundsDefensive": 32,
					"reboundsOffensive": 8,
					"reboundsPersonal": 40,
					"reboundsTeam": 4,
					"reboundsTeamDefensive": 2,
					"reboundsTeamOffensive": 2,
					"reboundsTotal": 44,
					"secondChancePointsAttempted": 8,
					"secondChancePointsMade": 5,
					"secondChancePointsPercentage": 0.625,
					"steals": 6,
					"threePointersAttempted": 30,
					"threePointersMade": 12,
					"threePointersPercentage": 0.4,
					"timeLeading": "PT5M00.00S",
					"timesTied": 2,
					"trueShootingAttempts": 95.0,
					"trueShootingPercentage": 0.553,
					"turnovers": 12,
					"turnoversTeam": 0,
					"turnoversTotal": 12,
					"twoPointersAttempted": 58,
					"twoPointersMade": 30,
					"twoPointersPercentage": 0.517
				}
			}
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "boxscore/boxscore_0022400001.json")
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

	boxscore, err := c.GetBoxScore(context.Background(), "0022400001")

	require.NoError(t, err)
	require.NotNil(t, boxscore)
	assert.Equal(t, 200, boxscore.Meta.Code)
	assert.Equal(t, "0022400001", boxscore.Game.GameID)
	assert.Equal(t, "Final", boxscore.Game.GameStatusText)
	assert.Equal(t, 110, boxscore.Game.HomeTeam.Score)
	assert.Equal(t, 105, boxscore.Game.AwayTeam.Score)
	assert.Equal(t, "Clippers", boxscore.Game.HomeTeam.TeamName)
	assert.Equal(t, "Lakers", boxscore.Game.AwayTeam.TeamName)
	assert.Len(t, boxscore.Game.HomeTeam.Players, 1)
	assert.Len(t, boxscore.Game.Officials, 1)
	assert.Equal(t, "Crypto.com Arena", boxscore.Game.Arena.ArenaName)
}

func TestGetBoxScore_InvalidJSON(t *testing.T) {
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

	_, err := c.GetBoxScore(context.Background(), "0022400001")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid JSON")
}

