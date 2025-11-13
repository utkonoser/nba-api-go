package leaders

import (
	"context"
	"fmt"
	"log/slog"
)

// AssistLeadersParams holds parameters for the AssistLeaders endpoint.
type AssistLeadersParams struct {
	LeagueId string
	PerModeSimple string
	PlayerOrTeam string
	Season string
	SeasonTypePlayoffs string
}

// GetAssistLeaders fetches data from the assistleaders endpoint.
func (c *Client) GetAssistLeaders(ctx context.Context, params AssistLeadersParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching assistleaders")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"PlayerOrTeam": params.PlayerOrTeam,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
	}

	resp, err := c.httpClient.SendRequest(ctx, "assistleaders", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch assistleaders",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch assistleaders: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from assistleaders endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal assistleaders response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched assistleaders",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
