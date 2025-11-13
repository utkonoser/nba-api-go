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

func TestGetTeamDetails(t *testing.T) {
	mockResponse := `{
		"resource": "teamdetails",
		"parameters": {},
		"resultSets": [
{
				"name": "TeamAwardsChampionships",
				"headers": ["YEARAWARDED", "OPPOSITETEAM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamAwardsConf",
				"headers": ["YEARAWARDED", "OPPOSITETEAM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamAwardsDiv",
				"headers": ["YEARAWARDED", "OPPOSITETEAM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamBackground",
				"headers": ["TEAM_ID", "ABBREVIATION", "NICKNAME", "YEARFOUNDED", "CITY"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamHistory",
				"headers": ["TEAM_ID", "CITY", "NICKNAME", "YEARFOUNDED", "YEARACTIVETILL"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamHof",
				"headers": ["PLAYERID", "PLAYER", "POSITION", "JERSEY", "SEASONSWITHTEAM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamRetired",
				"headers": ["PLAYERID", "PLAYER", "POSITION", "JERSEY", "SEASONSWITHTEAM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TeamSocialSites",
				"headers": ["ACCOUNTTYPE", "WEBSITE_LINK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamdetails")
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

	params := TeamDetailsParams{}

	response, err := c.GetTeamDetails(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamdetails", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
