package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerFantasyProfileBarGraphParams holds parameters for the PlayerFantasyProfileBarGraph endpoint.
type PlayerFantasyProfileBarGraphParams struct {
	PlayerId string
	Season string
	LeagueIdNullable string
	SeasonTypeAllStarNullable string
}

// GetPlayerFantasyProfileBarGraph fetches data from the playerfantasyprofilebargraph endpoint.
func (c *Client) GetPlayerFantasyProfileBarGraph(ctx context.Context, params PlayerFantasyProfileBarGraphParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerfantasyprofilebargraph")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"Season": params.Season,
		"LeagueID": params.LeagueIdNullable,
		"SeasonType": params.SeasonTypeAllStarNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerfantasyprofilebargraph", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerfantasyprofilebargraph",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerfantasyprofilebargraph: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerfantasyprofilebargraph endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerfantasyprofilebargraph response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerfantasyprofilebargraph",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
