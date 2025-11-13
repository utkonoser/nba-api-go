package leaders

import (
	"context"
	"fmt"
	"log/slog"
)

// DunkScoreLeadersParams holds parameters for the DunkScoreLeaders endpoint.
type DunkScoreLeadersParams struct {
	LeagueIdNullable string
	Season string
	SeasonTypeAllStar string
	PlayerIdNullable string
	TeamIdNullable string
	GameIdNullable string
}

// GetDunkScoreLeaders fetches data from the dunkscoreleaders endpoint.
func (c *Client) GetDunkScoreLeaders(ctx context.Context, params DunkScoreLeadersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching dunkscoreleaders")

	reqParams := map[string]string{
		"LeagueID": params.LeagueIdNullable,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"PlayerID": params.PlayerIdNullable,
		"TeamID": params.TeamIdNullable,
		"GameID": params.GameIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "dunkscoreleaders", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch dunkscoreleaders",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch dunkscoreleaders: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from dunkscoreleaders endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal dunkscoreleaders response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched dunkscoreleaders",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
