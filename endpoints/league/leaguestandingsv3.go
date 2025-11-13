package league

import (
	"context"
	"fmt"
	"log/slog"
)

// LeagueStandingsV3Params holds parameters for the LeagueStandingsV3 endpoint.
type LeagueStandingsV3Params struct {
	LeagueId string
	Season string
	SeasonType string
	SeasonNullable string
}

// GetLeagueStandingsV3 fetches data from the leaguestandingsv3 endpoint.
func (c *Client) GetLeagueStandingsV3(ctx context.Context, params LeagueStandingsV3Params) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching leaguestandingsv3")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"SeasonType": params.SeasonType,
		"SeasonYear": params.SeasonNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "leaguestandingsv3", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch leaguestandingsv3",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch leaguestandingsv3: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from leaguestandingsv3 endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal leaguestandingsv3 response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched leaguestandingsv3",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
