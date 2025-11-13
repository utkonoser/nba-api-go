package misc

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

func TestGetHomePageV2(t *testing.T) {
	mockResponse := `{
		"resource": "homepagev2",
		"parameters": {},
		"resultSets": [
{
				"name": "HomePageStat1",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "PTS"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat2",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "REB"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat3",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "AST"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat4",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "STL"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat5",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "FG_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat6",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "FT_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat7",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "FG3_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "HomePageStat8",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "BLK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "homepagev2")
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

	params := HomePageV2Params{}

	response, err := c.GetHomePageV2(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "homepagev2", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
