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

func TestGetTeamAndPlayersVsPlayers(t *testing.T) {
	mockResponse := `{
		"resource": "teamandplayersvsplayers",
		"parameters": {},
		"resultSets": [
{
				"name": "PlayersVsPlayers",
				"headers": ["GROUP_SET", "TITLE_DESCRIPTION", "DESCRIPTION", "MIN", "FGM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamPlayersVsPlayersOff",
				"headers": ["GROUP_SET", "TITLE_DESCRIPTION", "PLAYER_ID", "PLAYER_NAME", "MIN"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamPlayersVsPlayersOn",
				"headers": ["GROUP_SET", "TITLE_DESCRIPTION", "PLAYER_ID", "PLAYER_NAME", "MIN"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamVsPlayers",
				"headers": ["GROUP_SET", "TITLE_DESCRIPTION", "DESCRIPTION", "MIN", "FGM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamVsPlayersOff",
				"headers": ["GROUP_SET", "TITLE_DESCRIPTION", "DESCRIPTION", "MIN", "FGM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamandplayersvsplayers")
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

	params := TeamAndPlayersVsPlayersParams{}

	response, err := c.GetTeamAndPlayersVsPlayers(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamandplayersvsplayers", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
