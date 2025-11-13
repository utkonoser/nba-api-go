package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// MatchupsRollupParams holds parameters for the MatchupsRollup endpoint.
type MatchupsRollupParams struct {
	LeagueId string
	PerModeSimple string
	Season string
	SeasonTypePlayoffs string
	DefPlayerIdNullable string
	DefTeamIdNullable string
	OffPlayerIdNullable string
	OffTeamIdNullable string
}

// GetMatchupsRollup fetches data from the matchupsrollup endpoint.
func (c *Client) GetMatchupsRollup(ctx context.Context, params MatchupsRollupParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching matchupsrollup")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"DefPlayerID": params.DefPlayerIdNullable,
		"DefTeamID": params.DefTeamIdNullable,
		"OffPlayerID": params.OffPlayerIdNullable,
		"OffTeamID": params.OffTeamIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "matchupsrollup", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch matchupsrollup",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch matchupsrollup: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from matchupsrollup endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal matchupsrollup response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched matchupsrollup",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
