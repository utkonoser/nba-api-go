package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// GLAlumBoxScoreSimilarityScoreParams holds parameters for the GLAlumBoxScoreSimilarityScore endpoint.
type GLAlumBoxScoreSimilarityScoreParams struct {
	Person2Id string
	Person1Id string
	Person1LeagueId string
	Person1SeasonYear string
	Person1SeasonType string
	Person2LeagueId string
	Person2SeasonYear string
	Person2SeasonType string
}

// GetGLAlumBoxScoreSimilarityScore fetches data from the glalumboxscoresimilarityscore endpoint.
func (c *Client) GetGLAlumBoxScoreSimilarityScore(ctx context.Context, params GLAlumBoxScoreSimilarityScoreParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching glalumboxscoresimilarityscore")

	reqParams := map[string]string{
		"Person2Id": params.Person2Id,
		"Person1Id": params.Person1Id,
		"Person1LeagueId": params.Person1LeagueId,
		"Person1Season": params.Person1SeasonYear,
		"Person1SeasonType": params.Person1SeasonType,
		"Person2LeagueId": params.Person2LeagueId,
		"Person2Season": params.Person2SeasonYear,
		"Person2SeasonType": params.Person2SeasonType,
	}

	resp, err := c.httpClient.SendRequest(ctx, "glalumboxscoresimilarityscore", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch glalumboxscoresimilarityscore",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch glalumboxscoresimilarityscore: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from glalumboxscoresimilarityscore endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal glalumboxscoresimilarityscore response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched glalumboxscoresimilarityscore",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
