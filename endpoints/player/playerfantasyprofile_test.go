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

func TestGetPlayerFantasyProfile(t *testing.T) {
	mockResponse := `{
		"resource": "playerfantasyprofile",
		"parameters": {},
		"resultSets": [
{
				"name": "DaysRestModified",
				"headers": ["GROUP_SET", "GROUP_VALUE", "SEASON_YEAR", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "LastNGames",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Location",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Opponent",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Overall",
				"headers": ["GROUP_SET", "GROUP_VALUE", "GP", "W", "L"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerfantasyprofile")
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

	params := PlayerFantasyProfileParams{}

	response, err := c.GetPlayerFantasyProfile(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerfantasyprofile", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
