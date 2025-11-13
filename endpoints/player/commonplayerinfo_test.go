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

func TestGetCommonPlayerInfo(t *testing.T) {
	mockResponse := `{
		"resource": "commonplayerinfo",
		"parameters": {},
		"resultSets": [
{
				"name": "AvailableSeasons",
				"headers": ["SEASON_ID"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "CommonPlayerInfo",
				"headers": ["PERSON_ID", "FIRST_NAME", "LAST_NAME", "DISPLAY_FIRST_LAST", "DISPLAY_LAST_COMMA_FIRST"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PlayerHeadlineStats",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "TimeFrame", "PTS", "AST"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "commonplayerinfo")
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

	params := CommonPlayerInfoParams{}

	response, err := c.GetCommonPlayerInfo(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "commonplayerinfo", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
