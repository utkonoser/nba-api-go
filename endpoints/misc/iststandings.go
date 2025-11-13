package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// ISTStandingsParams holds parameters for the ISTStandings endpoint.
type ISTStandingsParams struct {
	LeagueId string
	Season string
	Section string
}

// GetISTStandings fetches data from the iststandings endpoint.
func (c *Client) GetISTStandings(ctx context.Context, params ISTStandingsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching iststandings")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"Season": params.Season,
		"Section": params.Section,
	}

	resp, err := c.httpClient.SendRequest(ctx, "iststandings", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch iststandings",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch iststandings: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from iststandings endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal iststandings response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched iststandings",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
