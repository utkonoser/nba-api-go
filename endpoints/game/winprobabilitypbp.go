package game

import (
	"context"
	"fmt"
	"log/slog"
)

// WinProbabilityPBPParams holds parameters for the WinProbabilityPBP endpoint.
type WinProbabilityPBPParams struct {
	GameId string
	RunType string
}

// GetWinProbabilityPBP fetches data from the winprobabilitypbp endpoint.
func (c *Client) GetWinProbabilityPBP(ctx context.Context, params WinProbabilityPBPParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching winprobabilitypbp")

	reqParams := map[string]string{
		"GameID": params.GameId,
		"RunType": params.RunType,
	}

	resp, err := c.httpClient.SendRequest(ctx, "winprobabilitypbp", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch winprobabilitypbp",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch winprobabilitypbp: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from winprobabilitypbp endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal winprobabilitypbp response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched winprobabilitypbp",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
