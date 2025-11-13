package game

import (
	"context"
	"fmt"
	"log/slog"
)

// GameRotationParams holds parameters for the GameRotation endpoint.
type GameRotationParams struct {
	GameId string
	LeagueId string
}

// GetGameRotation fetches data from the gamerotation endpoint.
func (c *Client) GetGameRotation(ctx context.Context, params GameRotationParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching gamerotation")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"LeagueID": params.LeagueId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "gamerotation", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch gamerotation",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch gamerotation: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from gamerotation endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal gamerotation response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched gamerotation",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
