package league

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

func TestGetLeagueLineupViz(t *testing.T) {
	mockResponse := `{
		"resource": "leaguelineupviz",
		"parameters": {},
		"resultSets": [
{
				"name": "LeagueLineupViz",
				"headers": ["GROUP_ID", "GROUP_NAME", "TEAM_ID", "TEAM_ABBREVIATION", "MIN"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "leaguelineupviz")
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

	params := LeagueLineupVizParams{}

	response, err := c.GetLeagueLineupViz(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "leaguelineupviz", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
