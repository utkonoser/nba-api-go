package franchise

import (
	"context"
	"fmt"
	"log/slog"
)

// FranchiseHistoryParams holds parameters for the FranchiseHistory endpoint.
type FranchiseHistoryParams struct {
	LeagueId string
}

// GetFranchiseHistory fetches data from the franchisehistory endpoint.
func (c *Client) GetFranchiseHistory(ctx context.Context, params FranchiseHistoryParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching franchisehistory")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "franchisehistory", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch franchisehistory",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch franchisehistory: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from franchisehistory endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal franchisehistory response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched franchisehistory",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
