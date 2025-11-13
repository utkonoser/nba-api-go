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

func TestGetTeamDashLineups(t *testing.T) {
	mockResponse := `{
		"resource": "teamdashlineups",
		"parameters": {},
		"resultSets": [
{
				"name": "Lineups",
				"headers": ["GROUP_SET", "GROUP_ID", "GROUP_NAME", "GP", "W"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "Overall",
				"headers": ["GROUP_SET", "GROUP_VALUE", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamdashlineups")
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

	params := TeamDashLineupsParams{}

	response, err := c.GetTeamDashLineups(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamdashlineups", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
