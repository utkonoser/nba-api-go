package boxscore

import (
	"context"
	"fmt"
	"log/slog"
)

// BoxScoreTraditionalV3Params holds parameters for the BoxScoreTraditionalV3 endpoint.
type BoxScoreTraditionalV3Params struct {
	GameId string
	EndPeriod string
	EndRange string
	RangeType string
	StartPeriod string
	StartRange string
}

// GetBoxScoreTraditionalV3 fetches data from the boxscoretraditionalv3 endpoint.
func (c *Client) GetBoxScoreTraditionalV3(ctx context.Context, params BoxScoreTraditionalV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching boxscoretraditionalv3")

	reqParams := map[string]string{
		"GameID": params.GameId,
	}
	
	// Add optional parameters only if they are not empty
	if params.EndPeriod != "" {
		reqParams["EndPeriod"] = params.EndPeriod
	}
	if params.EndRange != "" {
		reqParams["EndRange"] = params.EndRange
	}
	if params.RangeType != "" {
		reqParams["RangeType"] = params.RangeType
	}
	if params.StartPeriod != "" {
		reqParams["StartPeriod"] = params.StartPeriod
	}
	if params.StartRange != "" {
		reqParams["StartRange"] = params.StartRange
	}

	resp, err := c.httpClient.SendRequest(ctx, "boxscoretraditionalv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch boxscoretraditionalv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch boxscoretraditionalv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from boxscoretraditionalv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal boxscoretraditionalv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched boxscoretraditionalv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
