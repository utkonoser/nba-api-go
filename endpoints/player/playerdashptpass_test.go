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

func TestGetPlayerDashPtPass(t *testing.T) {
	mockResponse := `{
		"resource": "playerdashptpass",
		"parameters": {},
		"resultSets": [
{
				"name": "PassesMade",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "TEAM_NAME", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PassesReceived",
				"headers": ["PLAYER_ID", "PLAYER_NAME_LAST_FIRST", "TEAM_NAME", "TEAM_ID", "TEAM_ABBREVIATION"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerdashptpass")
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

	params := PlayerDashPtPassParams{}

	response, err := c.GetPlayerDashPtPass(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerdashptpass", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
