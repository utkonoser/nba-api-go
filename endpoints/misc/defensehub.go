package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// DefenseHubParams holds parameters for the DefenseHub endpoint.
type DefenseHubParams struct {
	GameScopeDetailed string
	LeagueId string
	PlayerOrTeam string
	PlayerScope string
	Season string
	SeasonTypePlayoffs string
}

// GetDefenseHub fetches data from the defensehub endpoint.
func (c *Client) GetDefenseHub(ctx context.Context, params DefenseHubParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching defensehub")

	reqParams := map[string]string{
		"GameScope": params.GameScopeDetailed,
		"LeagueID": params.LeagueId,
		"PlayerOrTeam": params.PlayerOrTeam,
		"PlayerScope": params.PlayerScope,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
	}

	resp, err := c.httpClient.SendRequest(ctx, "defensehub", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch defensehub",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch defensehub: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from defensehub endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal defensehub response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched defensehub",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
