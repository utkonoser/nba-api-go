package franchise

import (
	"context"
	"fmt"
	"log/slog"
)

// FranchisePlayersParams holds parameters for the FranchisePlayers endpoint.
type FranchisePlayersParams struct {
	TeamId string
	LeagueId string
	PerModeDetailed string
	SeasonTypeAllStar string
}

// GetFranchisePlayers fetches data from the franchiseplayers endpoint.
func (c *Client) GetFranchisePlayers(ctx context.Context, params FranchisePlayersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching franchiseplayers")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeDetailed,
		"SeasonType": params.SeasonTypeAllStar,
	}

	resp, err := c.httpClient.SendRequest(ctx, "franchiseplayers", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch franchiseplayers",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch franchiseplayers: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from franchiseplayers endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal franchiseplayers response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched franchiseplayers",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
