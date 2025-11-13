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

func TestGetGLAlumBoxScoreSimilarityScore(t *testing.T) {
	mockResponse := `{
		"resource": "glalumboxscoresimilarityscore",
		"parameters": {},
		"resultSets": [
{
				"name": "GLeagueAlumBoxScoreSimilarityScores",
				"headers": ["PERSON_2_ID", "PERSON_2", "TEAM_ID", "SIMILARITY_SCORE"],
				"rowSet": [[1, "Test", "Data", 100, 50]]
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "glalumboxscoresimilarityscore")
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

	params := GLAlumBoxScoreSimilarityScoreParams{}

	response, err := c.GetGLAlumBoxScoreSimilarityScore(context.Background(), params)

	require.NoError(t, err)
	require.NotNil(t, response)
	assert.Equal(t, "glalumboxscoresimilarityscore", response.Resource)
	// Note: ResultSets may be empty for some endpoints
	assert.NotNil(t, response.ResultSets)
}
