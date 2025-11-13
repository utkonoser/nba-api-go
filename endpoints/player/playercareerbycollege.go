package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerCareerByCollegeParams holds parameters for the PlayerCareerByCollege endpoint.
type PlayerCareerByCollegeParams struct {
	College string
	LeagueId string
	PerModeSimple string
	SeasonTypeAllStar string
	SeasonNullable string
}

// GetPlayerCareerByCollege fetches data from the playercareerbycollege endpoint.
func (c *Client) GetPlayerCareerByCollege(ctx context.Context, params PlayerCareerByCollegeParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playercareerbycollege")

	reqParams := map[string]string{
		"College": params.College,
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"SeasonType": params.SeasonTypeAllStar,
		"Season": params.SeasonNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playercareerbycollege", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playercareerbycollege",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playercareerbycollege: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playercareerbycollege endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playercareerbycollege response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playercareerbycollege",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
