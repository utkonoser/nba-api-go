package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerCareerByCollegeRollupParams holds parameters for the PlayerCareerByCollegeRollup endpoint.
type PlayerCareerByCollegeRollupParams struct {
	LeagueId string
	PerModeSimple string
	SeasonTypeAllStar string
	SeasonNullable string
}

// GetPlayerCareerByCollegeRollup fetches data from the playercareerbycollegerollup endpoint.
func (c *Client) GetPlayerCareerByCollegeRollup(ctx context.Context, params PlayerCareerByCollegeRollupParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playercareerbycollegerollup")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"SeasonType": params.SeasonTypeAllStar,
		"Season": params.SeasonNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playercareerbycollegerollup", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playercareerbycollegerollup",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playercareerbycollegerollup: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playercareerbycollegerollup endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playercareerbycollegerollup response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playercareerbycollegerollup",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
