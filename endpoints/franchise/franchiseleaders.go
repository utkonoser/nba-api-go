package franchise

import (
	"context"
	"fmt"
	"log/slog"
)

// FranchiseLeadersParams holds parameters for the FranchiseLeaders endpoint.
type FranchiseLeadersParams struct {
	TeamId string
	LeagueIdNullable string
}

// GetFranchiseLeaders fetches data from the franchiseleaders endpoint.
func (c *Client) GetFranchiseLeaders(ctx context.Context, params FranchiseLeadersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching franchiseleaders")

	reqParams := map[string]string{
		"TeamID": params.TeamId,
	}
	
	// Add optional parameters only if they are not empty
	if params.LeagueIdNullable != "" {
		reqParams["LeagueID"] = params.LeagueIdNullable
	}

	resp, err := c.httpClient.SendRequest(ctx, "franchiseleaders", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch franchiseleaders",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch franchiseleaders: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from franchiseleaders endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal franchiseleaders response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched franchiseleaders",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
