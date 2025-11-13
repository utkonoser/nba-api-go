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

func TestGetDefenseHub(t *testing.T) {
	mockResponse := `{
		"resource": "defensehub",
		"parameters": {},
		"resultSets": [
{
				"name": "DefenseHubStat1",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "DREB"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat10",
				"headers": [],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat2",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "STL"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat3",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "BLK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat4",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "TM_DEF_RATING"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat5",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "OVERALL_PM"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat6",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "THREEP_DFGPCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat7",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "TWOP_DFGPCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat8",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "FIFETEENF_DFGPCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DefenseHubStat9",
				"headers": ["RANK", "TEAM_ID", "TEAM_ABBREVIATION", "TEAM_NAME", "DEF_RIM_PCT"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "defensehub")
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

	params := DefenseHubParams{}

	response, err := c.GetDefenseHub(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "defensehub", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
