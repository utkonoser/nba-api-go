package game

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

func TestGetWinProbabilityPBP(t *testing.T) {
	mockResponse := `{
		"resource": "winprobabilitypbp",
		"parameters": {},
		"resultSets": [
{
				"name": "GameInfo",
				"headers": ["GAME_ID", "GAME_DATE", "HOME_TEAM_ID", "HOME_TEAM_ABR", "HOME_TEAM_PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "WinProbPBP",
				"headers": ["GAME_ID", "EVENT_NUM", "HOME_PCT", "VISITOR_PCT", "HOME_PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "winprobabilitypbp")
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

	params := WinProbabilityPBPParams{}

	response, err := c.GetWinProbabilityPBP(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "winprobabilitypbp", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
