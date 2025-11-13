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

func TestGetPlayerVsPlayer(t *testing.T) {
	mockResponse := `{
		"resource": "playervsplayer",
		"parameters": {},
		"resultSets": [
{
				"name": "OnOffCourt",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "VS_PLAYER_ID", "VS_PLAYER_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Overall",
				"headers": ["GROUP_SET", "GROUP_VALUE", "PLAYER_ID", "PLAYER_NAME", "GP"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PlayerInfo",
				"headers": ["PERSON_ID", "FIRST_NAME", "LAST_NAME", "DISPLAY_FIRST_LAST", "DISPLAY_LAST_COMMA_FIRST"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotAreaOffCourt",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "VS_PLAYER_ID", "VS_PLAYER_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotAreaOnCourt",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "VS_PLAYER_ID", "VS_PLAYER_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotAreaOverall",
				"headers": ["GROUP_SET", "GROUP_VALUE", "PLAYER_ID", "PLAYER_NAME", "FGM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotDistanceOffCourt",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "VS_PLAYER_ID", "VS_PLAYER_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotDistanceOnCourt",
				"headers": ["GROUP_SET", "PLAYER_ID", "PLAYER_NAME", "VS_PLAYER_ID", "VS_PLAYER_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "ShotDistanceOverall",
				"headers": ["GROUP_SET", "GROUP_VALUE", "PLAYER_ID", "PLAYER_NAME", "FGM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "VsPlayerInfo",
				"headers": ["PERSON_ID", "FIRST_NAME", "LAST_NAME", "DISPLAY_FIRST_LAST", "DISPLAY_LAST_COMMA_FIRST"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playervsplayer")
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

	params := PlayerVsPlayerParams{}

	response, err := c.GetPlayerVsPlayer(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playervsplayer", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
