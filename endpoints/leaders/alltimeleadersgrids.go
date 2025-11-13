package leaders

import (
	"context"
	"fmt"
	"log/slog"
)

// AllTimeLeadersGridsParams holds parameters for the AllTimeLeadersGrids endpoint.
type AllTimeLeadersGridsParams struct {
	LeagueId string
	PerModeSimple string
	SeasonType string
	Topx string
}

// GetAllTimeLeadersGrids fetches data from the alltimeleadersgrids endpoint.
func (c *Client) GetAllTimeLeadersGrids(ctx context.Context, params AllTimeLeadersGridsParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching alltimeleadersgrids")

	reqParams := map[string]string{
		"LeagueID": params.LeagueId,
		"PerMode": params.PerModeSimple,
		"SeasonType": params.SeasonType,
		"TopX": params.Topx,
	}

	resp, err := c.httpClient.SendRequest(ctx, "alltimeleadersgrids", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch alltimeleadersgrids",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch alltimeleadersgrids: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from alltimeleadersgrids endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal alltimeleadersgrids response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched alltimeleadersgrids",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
