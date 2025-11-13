package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerProfileV2Params holds parameters for the PlayerProfileV2 endpoint.
type PlayerProfileV2Params struct {
	PlayerId string
	PerMode36 string
	LeagueIdNullable string
}

// GetPlayerProfileV2 fetches data from the playerprofilev2 endpoint.
func (c *Client) GetPlayerProfileV2(ctx context.Context, params PlayerProfileV2Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerprofilev2")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"PerMode": params.PerMode36,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerprofilev2", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerprofilev2",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerprofilev2: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerprofilev2 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerprofilev2 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerprofilev2",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
