package misc

import (
	"context"
	"fmt"
	"log/slog"
)

// SynergyPlayTypesParams holds parameters for the SynergyPlayTypes endpoint.
type SynergyPlayTypesParams struct {
	LeagueId string
	PerModeSimple string
	PlayerOrTeamAbbreviation string
	SeasonTypeAllStar string
	Season string
	PlayTypeNullable string
	TypeGroupingNullable string
}

// GetSynergyPlayTypes fetches data from the synergyplaytypes endpoint.
func (c *Client) GetSynergyPlayTypes(ctx context.Context, params SynergyPlayTypesParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching synergyplaytypes")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"PlayerOrTeam": params.PlayerOrTeamAbbreviation,
		"SeasonType": params.SeasonTypeAllStar,
		"SeasonYear": params.Season,
		"PlayType": params.PlayTypeNullable,
		"TypeGrouping": params.TypeGroupingNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "synergyplaytypes", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch synergyplaytypes",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch synergyplaytypes: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from synergyplaytypes endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal synergyplaytypes response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched synergyplaytypes",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
