package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerGameLogParams holds parameters for the PlayerGameLog endpoint.
type PlayerGameLogParams struct {
	PlayerId string
	Season string
	SeasonTypeAllStar string
	DateFromNullable string
	DateToNullable string
	LeagueIdNullable string
}

// GetPlayerGameLog fetches data from the playergamelog endpoint.
func (c *Client) GetPlayerGameLog(ctx context.Context, params PlayerGameLogParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playergamelog")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"Season": params.Season,
		"SeasonType": params.SeasonTypeAllStar,
		"DateFrom": params.DateFromNullable,
		"DateTo": params.DateToNullable,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playergamelog", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playergamelog",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playergamelog: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playergamelog endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playergamelog response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playergamelog",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
