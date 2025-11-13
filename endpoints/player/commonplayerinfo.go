package player

import (
	"context"
	"fmt"
	"log/slog"
)

// CommonPlayerInfoParams holds parameters for the CommonPlayerInfo endpoint.
type CommonPlayerInfoParams struct {
	PlayerId string
	LeagueIdNullable string
}

// GetCommonPlayerInfo fetches data from the commonplayerinfo endpoint.
func (c *Client) GetCommonPlayerInfo(ctx context.Context, params CommonPlayerInfoParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching commonplayerinfo")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "commonplayerinfo", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch commonplayerinfo",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch commonplayerinfo: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from commonplayerinfo endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal commonplayerinfo response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched commonplayerinfo",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
