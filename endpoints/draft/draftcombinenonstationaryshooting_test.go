package draft

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

func TestGetDraftCombineNonStationaryShooting(t *testing.T) {
	mockResponse := `{
		"resource": "draftcombinenonstationaryshooting",
		"parameters": {},
		"resultSets": [
{
				"name": "Results",
				"headers": ["TEMP_PLAYER_ID", "PLAYER_ID", "FIRST_NAME", "LAST_NAME", "PLAYER_NAME"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "draftcombinenonstationaryshooting")
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

	params := DraftCombineNonStationaryShootingParams{}

	response, err := c.GetDraftCombineNonStationaryShooting(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "draftcombinenonstationaryshooting", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
