package player

import (
	"context"
	"fmt"
	"log/slog"
)

// CommonAllPlayersParams holds parameters for the CommonAllPlayers endpoint.
type CommonAllPlayersParams struct {
	IsOnlyCurrentSeason string
	LeagueId string
	Season string
}

// GetCommonAllPlayers fetches data from the commonallplayers endpoint.
func (c *Client) GetCommonAllPlayers(ctx context.Context, params CommonAllPlayersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching commonallplayers")

	reqParams := map[string]string{
		"IsOnlyCurrentSeason": params.IsOnlyCurrentSeason,
		"LeagueID": params.LeagueId,
		"Season": params.Season,
	}

	resp, err := c.httpClient.SendRequest(ctx, "commonallplayers", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch commonallplayers",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch commonallplayers: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from commonallplayers endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal commonallplayers response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched commonallplayers",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
