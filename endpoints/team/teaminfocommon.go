package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamInfoCommonParams holds parameters for the TeamInfoCommon endpoint.
type TeamInfoCommonParams struct {
	TeamId string
	LeagueId string
	SeasonNullable string
	SeasonTypeNullable string
}

// GetTeamInfoCommon fetches data from the teaminfocommon endpoint.
func (c *Client) GetTeamInfoCommon(ctx context.Context, params TeamInfoCommonParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teaminfocommon")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LeagueID": params.LeagueId,
		"Season": params.SeasonNullable,
		"SeasonType": params.SeasonTypeNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teaminfocommon", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teaminfocommon",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teaminfocommon: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teaminfocommon endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teaminfocommon response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teaminfocommon",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
