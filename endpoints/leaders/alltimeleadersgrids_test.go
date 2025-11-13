package leaders

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

func TestGetAllTimeLeadersGrids(t *testing.T) {
	mockResponse := `{
		"resource": "alltimeleadersgrids",
		"parameters": {},
		"resultSets": [
{
				"name": "ASTLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "AST", "AST_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "BLKLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "BLK", "BLK_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "DREBLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "DREB", "DREB_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FG3ALeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FG3A", "FG3A_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FG3MLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FG3M", "FG3M_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FG3_PCTLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FG3_PCT", "FG3_PCT_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FGALeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FGA", "FGA_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FGMLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FGM", "FGM_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FG_PCTLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FG_PCT", "FG_PCT_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FTALeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FTA", "FTA_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FTMLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FTM", "FTM_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "FT_PCTLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "FT_PCT", "FT_PCT_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "GPLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "GP", "GP_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "OREBLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "OREB", "OREB_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PFLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "PF", "PF_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "PTSLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "PTS", "PTS_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "REBLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "REB", "REB_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "STLLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "STL", "STL_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			},
{
				"name": "TOVLeaders",
				"headers": ["PLAYER_ID", "PLAYER_NAME", "TOV", "TOV_RANK"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "alltimeleadersgrids")
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

	params := AllTimeLeadersGridsParams{}

	response, err := c.GetAllTimeLeadersGrids(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "alltimeleadersgrids", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
