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

func TestGetTeamDashPtPass(t *testing.T) {
	mockResponse := `{
		"resource": "teamdashptpass",
		"parameters": {},
		"resultSets": [
{
				"name": "PassesMade",
				"headers": ["TEAM_ID", "TEAM_NAME", "PASS_TYPE", "G", "PASS_FROM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PassesReceived",
				"headers": ["TEAM_ID", "TEAM_NAME", "PASS_TYPE", "G", "PASS_TO"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "teamdashptpass")
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

	params := TeamDashPtPassParams{}

	response, err := c.GetTeamDashPtPass(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "teamdashptpass", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
