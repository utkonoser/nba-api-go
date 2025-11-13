package playoff

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

func TestGetPlayoffPicture(t *testing.T) {
	mockResponse := `{
		"resource": "playoffpicture",
		"parameters": {},
		"resultSets": [
{
				"name": "EastConfPlayoffPicture",
				"headers": ["CONFERENCE", "HIGH_SEED_RANK", "HIGH_SEED_TEAM", "HIGH_SEED_TEAM_ID", "LOW_SEED_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "EastConfRemainingGames",
				"headers": ["TEAM", "TEAM_ID", "REMAINING_G", "REMAINING_HOME_G", "REMAINING_AWAY_G"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "EastConfStandings",
				"headers": ["CONFERENCE", "RANK", "TEAM", "TEAM_SLUG", "TEAM_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "WestConfPlayoffPicture",
				"headers": ["CONFERENCE", "HIGH_SEED_RANK", "HIGH_SEED_TEAM", "HIGH_SEED_TEAM_ID", "LOW_SEED_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "WestConfRemainingGames",
				"headers": ["TEAM", "TEAM_ID", "REMAINING_G", "REMAINING_HOME_G", "REMAINING_AWAY_G"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "WestConfStandings",
				"headers": ["CONFERENCE", "RANK", "TEAM", "TEAM_SLUG", "TEAM_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playoffpicture")
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

	params := PlayoffPictureParams{}

	response, err := c.GetPlayoffPicture(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playoffpicture", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
