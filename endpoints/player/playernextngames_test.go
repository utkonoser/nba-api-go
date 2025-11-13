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

func TestGetPlayerNextNGames(t *testing.T) {
	mockResponse := `{
		"resource": "playernextngames",
		"parameters": {},
		"resultSets": [
{
				"name": "NextNGames",
				"headers": ["GAME_ID", "GAME_DATE", "HOME_TEAM_ID", "VISITOR_TEAM_ID", "HOME_TEAM_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playernextngames")
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

	params := PlayerNextNGamesParams{}

	response, err := c.GetPlayerNextNGames(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playernextngames", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
