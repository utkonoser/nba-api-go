package player

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayerFantasyProfileParams holds parameters for the PlayerFantasyProfile endpoint.
type PlayerFantasyProfileParams struct {
	PlayerId string
	MeasureTypeBase string
	PaceAdjustNo string
	PerMode36 string
	PlusMinusNo string
	RankNo string
	Season string
	SeasonTypePlayoffs string
	LeagueIdNullable string
}

// GetPlayerFantasyProfile fetches data from the playerfantasyprofile endpoint.
func (c *Client) GetPlayerFantasyProfile(ctx context.Context, params PlayerFantasyProfileParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerfantasyprofile")

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"MeasureType": params.MeasureTypeBase,
		"PaceAdjust": params.PaceAdjustNo,
		"PerMode": params.PerMode36,
		"PlusMinus": params.PlusMinusNo,
		"Rank": params.RankNo,
		"Season": params.Season,
		"SeasonType": params.SeasonTypePlayoffs,
		"LeagueID": params.LeagueIdNullable,
	}

	resp, err := c.httpClient.SendRequest(ctx, "playerfantasyprofile", reqParams)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch playerfantasyprofile",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch playerfantasyprofile: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from playerfantasyprofile endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var statsResp StatsResponse
	if err := resp.GetJSON(&statsResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal playerfantasyprofile response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched playerfantasyprofile",
		slog.Int("result_sets_count", len(statsResp.ResultSets)))

	return &statsResp, nil
}
