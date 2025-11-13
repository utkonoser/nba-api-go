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

func TestGetTeamPlayerOnOffDetails(t *testing.T) {
	mockResponse := `{
		"resource": "teamplayeronoffdetails",
		"parameters": {},
		"resultSets": [
{
				"name": "OverallTeamPlayerOnOffDetails",
				"headers": ["GROUP_SET", "GROUP_VALUE", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PlayersOffCourtTeamPlayerOnOffDetails",
				"headers": ["GROUP_SET", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "VS_PLAYER_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PlayersOnCourtTeamPlayerOnOffDetails",
				"headers": ["GROUP_SET", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "VS_PLAYER_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamplayeronoffdetails")
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

	params := TeamPlayerOnOffDetailsParams{}

	response, err := c.GetTeamPlayerOnOffDetails(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamplayeronoffdetails", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
