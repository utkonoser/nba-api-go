package tracking

import (
	"context"
	"fmt"
	"log/slog"
)

// CumeStatsPlayerParams holds parameters for the CumeStatsPlayer endpoint.
type CumeStatsPlayerParams struct {
	PlayerId string
	GameIds string
	LeagueId string
	Season string
	SeasonTypeAllStar string
}

// GetCumeStatsPlayer fetches data from the cumestatsplayer endpoint.
func (c *Client) GetCumeStatsPlayer(ctx context.Context, params CumeStatsPlayerParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching cumestatsplayer")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"GameIDs": params.GameIds,
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
	}

	resp, err := c.httpClient.SendRequest(ctx, "cumestatsplayer", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch cumestatsplayer",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch cumestatsplayer: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from cumestatsplayer endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal cumestatsplayer response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched cumestatsplayer",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
