package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreFourFactorsV3Params holds parameters for the BoxScoreFourFactorsV3 endpoint.
type BoxScoreFourFactorsV3Params struct {
	GameId string
	EndPeriod string
	EndRange string
	RangeType string
	StartPeriod string
	StartRange string
}

// GetBoxScoreFourFactorsV3 fetches data from the boxscorefourfactorsv3 endpoint.
func (c *Client) GetBoxScoreFourFactorsV3(ctx context.Context, params BoxScoreFourFactorsV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscorefourfactorsv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"EndPeriod": params.EndPeriod,
		"EndRange": params.EndRange,
		"RangeType": params.RangeType,
		"StartPeriod": params.StartPeriod,
		"StartRange": params.StartRange,
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscorefourfactorsv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscorefourfactorsv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscorefourfactorsv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscorefourfactorsv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscorefourfactorsv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscorefourfactorsv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
