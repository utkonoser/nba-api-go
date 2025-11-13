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

func TestGetPlayerAwards(t *testing.T) {
	mockResponse := `{
		"resource": "playerawards",
		"parameters": {},
		"resultSets": [
			{
				"name": "PlayerAwards",
				"headers": ["PERSON_ID", "FIRST_NAME", "LAST_NAME", "TEAM", "DESCRIPTION"],
				"rowSet": [[2544, "LeBron", "James", "LAL", "NBA Championship"]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "playerawards")
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

	params := PlayerAwardsParams{
		PlayerId: "2544",
	}

	response, err := c.GetPlayerAwards(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "playerawards", response.Resource)
	assert.NotNil(t, response.ResultSets)
}

