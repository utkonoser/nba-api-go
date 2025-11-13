package team

import (
	"context"
	"fmt"
	"log/slog"
)

// TeamDetailsParams holds parameters for the TeamDetails endpoint.
type TeamDetailsParams struct {
	TeamId string
}

// GetTeamDetails fetches data from the teamdetails endpoint.
func (c *Client) GetTeamDetails(ctx context.Context, params TeamDetailsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching teamdetails")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
	}

	resp, err := c.httpClient.SendRequest(ctx, "teamdetails", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch teamdetails",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch teamdetails: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from teamdetails endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal teamdetails response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched teamdetails",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
