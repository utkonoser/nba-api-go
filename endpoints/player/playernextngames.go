package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerNextNGamesParams holds parameters for the PlayerNextNGames endpoint.
type PlayerNextNGamesParams struct {
	PlayerId string
	NumberOfGames string
	SeasonAll string
	SeasonTypeAllStar string
	LeagueIdNullable string
}

// GetPlayerNextNGames fetches data from the playernextngames endpoint.
func (c *Client) GetPlayerNextNGames(ctx context.Context, params PlayerNextNGamesParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playernextngames")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"NumberOfGames": params.NumberOfGames,
		"Season": params.SeasonAll,
		"SeasonType": params.SeasonTypeAllStar,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playernextngames", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playernextngames",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playernextngames: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playernextngames endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playernextngames response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playernextngames",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
