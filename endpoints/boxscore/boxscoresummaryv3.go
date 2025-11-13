package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreSummaryV3Params holds parameters for the BoxScoreSummaryV3 endpoint.
type BoxScoreSummaryV3Params struct {
	GameId string
}

// GetBoxScoreSummaryV3 fetches data from the boxscoresummaryv3 endpoint.
func (c *Client) GetBoxScoreSummaryV3(ctx context.Context, params BoxScoreSummaryV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscoresummaryv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscoresummaryv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscoresummaryv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscoresummaryv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscoresummaryv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscoresummaryv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscoresummaryv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
