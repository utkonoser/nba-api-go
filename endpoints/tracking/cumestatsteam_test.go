package tracking

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

func TestGetCumeStatsTeam(t *testing.T) {
	mockResponse := `{
		"resource": "cumestatsteam",
		"parameters": {},
		"resultSets": [
{
				"name": "GameByGameStats",
				"headers": ["JERSEY_NUM", "PLAYER", "PERSON_ID", "TEAM_ID", "GP"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TotalTeamStats",
				"headers": ["CITY", "NICKNAME", "TEAM_ID", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "cumestatsteam")
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

	params := CumeStatsTeamParams{}

	response, err := c.GetCumeStatsTeam(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "cumestatsteam", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
