package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// InfographicFanDuelPlayerParams holds parameters for the InfographicFanDuelPlayer endpoint.
type InfographicFanDuelPlayerParams struct {
	GameId string
}

// GetInfographicFanDuelPlayer fetches data from the infographicfanduelplayer endpoint.
func (c *Client) GetInfographicFanDuelPlayer(ctx context.Context, params InfographicFanDuelPlayerParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching infographicfanduelplayer")

	reqParams := map[string]string{
		"GameID": params.GameId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "infographicfanduelplayer", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch infographicfanduelplayer",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch infographicfanduelplayer: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from infographicfanduelplayer endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal infographicfanduelplayer response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched infographicfanduelplayer",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
