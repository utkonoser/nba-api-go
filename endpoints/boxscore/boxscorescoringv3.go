package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreScoringV3Params holds parameters for the BoxScoreScoringV3 endpoint.
type BoxScoreScoringV3Params struct {
	GameId string
	EndPeriod string
	EndRange string
	RangeType string
	StartPeriod string
	StartRange string
}

// GetBoxScoreScoringV3 fetches data from the boxscorescoringv3 endpoint.
func (c *Client) GetBoxScoreScoringV3(ctx context.Context, params BoxScoreScoringV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscorescoringv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"EndPeriod": params.EndPeriod,
		"EndRange": params.EndRange,
		"RangeType": params.RangeType,
		"StartPeriod": params.StartPeriod,
		"StartRange": params.StartRange,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscorescoringv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscorescoringv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscorescoringv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscorescoringv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscorescoringv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscorescoringv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
