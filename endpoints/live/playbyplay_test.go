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

func TestGetPlayByPlay(t *testing.T) {
	mockResponse := `{
		"meta": {
			"version": 1,
			"code": 200,
			"request": "",
			"time": "2024-01-01T00:00:00Z"
		},
		"game": {
			"gameId": "0022400001",
			"actions": [
				{
					"actionNumber": 1,
					"clock": "PT12M00.00S",
					"timeActual": "2024-01-01T19:30:00Z",
					"period": 1,
					"periodType": "REGULAR",
					"teamId": 1610612746,
					"teamTricode": "LAC",
					"actionType": "jumpball",
					"subType": "",
					"descriptor": "",
					"qualifiers": [],
					"personId": 1629029,
					"x": null,
					"y": null,
					"possession": 1610612746,
					"scoreHome": "0",
					"scoreAway": "0",
					"edited": "2024-01-01T19:30:00Z",
					"orderNumber": 1,
					"xLegacy": null,
					"yLegacy": null,
					"isFieldGoal": 0,
					"shotDistance": null,
					"shotResult": "",
					"pointsTotal": 0,
					"description": "Jump Ball",
					"personIdsFilter": [1629029],
					"assistPersonId": 0,
					"assistPlayerNameInitial": "",
					"assistTotal": 0,
					"officialId": 0,
					"foulDrawnPersonId": 0,
					"foulPersonalTotal": 0,
					"foulTechnicalTotal": 0,
					"shotActionNumber": 0,
					"reboundTotal": 0,
					"reboundDefensiveTotal": 0,
					"reboundOffensiveTotal": 0,
					"turnoverTotal": 0,
					"stealPersonId": 0,
					"value": "",
					"jumpBallWonPersonId": 1629029,
					"jumpBallLostPersonId": 2544,
					"shotWasBlocked": 0,
					"blockPersonId": 0,
					"playerName": "Player One",
					"playerNameI": "P. One"
				},
				{
					"actionNumber": 2,
					"clock": "PT11M45.00S",
					"timeActual": "2024-01-01T19:30:15Z",
					"period": 1,
					"periodType": "REGULAR",
					"teamId": 1610612746,
					"teamTricode": "LAC",
					"actionType": "2pt",
					"subType": "layup",
					"descriptor": "MADE",
					"qualifiers": ["fastbreak"],
					"personId": 1629029,
					"x": 5.0,
					"y": 25.0,
					"possession": 1610612746,
					"scoreHome": "2",
					"scoreAway": "0",
					"edited": "2024-01-01T19:30:15Z",
					"orderNumber": 2,
					"xLegacy": 50,
					"yLegacy": 250,
					"isFieldGoal": 1,
					"shotDistance": 5.2,
					"shotResult": "Made",
					"pointsTotal": 2,
					"description": "Player One 5' Layup (2 PTS)",
					"personIdsFilter": [1629029],
					"assistPersonId": 203999,
					"assistPlayerNameInitial": "A. Player",
					"assistTotal": 1,
					"officialId": 0,
					"foulDrawnPersonId": 0,
					"foulPersonalTotal": 0,
					"foulTechnicalTotal": 0,
					"shotActionNumber": 2,
					"reboundTotal": 0,
					"reboundDefensiveTotal": 0,
					"reboundOffensiveTotal": 0,
					"turnoverTotal": 0,
					"stealPersonId": 0,
					"value": "",
					"jumpBallWonPersonId": 0,
					"jumpBallLostPersonId": 0,
					"shotWasBlocked": 0,
					"blockPersonId": 0,
					"playerName": "Player One",
					"playerNameI": "P. One"
				}
			]
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playbyplay/playbyplay_0022400001.json")
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

	pbp, err := c.GetPlayByPlay(context.Background(), "0022400001")

	require.NoError(t, err)
	require.NotNil(t, pbp)
	assert.Equal(t, 200, pbp.Meta.Code)
	assert.Equal(t, "0022400001", pbp.Game.GameID)
	assert.Len(t, pbp.Game.Actions, 2)

	action := pbp.Game.Actions[0]
	assert.Equal(t, 1, action.ActionNumber)
	assert.Equal(t, "jumpball", action.ActionType)
	assert.Equal(t, "Player One", action.PlayerName)

	action2 := pbp.Game.Actions[1]
	assert.Equal(t, 2, action2.ActionNumber)
	assert.Equal(t, "2pt", action2.ActionType)
	assert.Equal(t, "MADE", action2.Descriptor)
}

func TestGetPlayByPlay_InvalidJSON(t *testing.T) {
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

	_, err := c.GetPlayByPlay(context.Background(), "0022400001")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid JSON")
}

