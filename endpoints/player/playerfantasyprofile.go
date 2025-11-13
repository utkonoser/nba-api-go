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
	SeasonType string
	LeagueIdNullable string
}

// GetPlayerFantasyProfile fetches data from the playerfantasyprofile endpoint.
func (c *Client) GetPlayerFantasyProfile(ctx context.Context, params PlayerFantasyProfileParams) (*StatsResponse, error) {
	c.logger.InfoContext(ctx, "Fetching playerfantasyprofile")

	// Set defaults for required parameters
	measureType := params.MeasureTypeBase
	if measureType == "" {
		measureType = "Base"
	}
	paceAdjust := params.PaceAdjustNo
	if paceAdjust == "" {
		paceAdjust = "N"
	}
	perMode := params.PerMode36
	if perMode == "" {
		perMode = "Per36"
	}
	plusMinus := params.PlusMinusNo
	if plusMinus == "" {
		plusMinus = "N"
	}
	rank := params.RankNo
	if rank == "" {
		rank = "N"
	}
	season := params.Season
	if season == "" {
		season = "2023-24"
	}
	seasonType := params.SeasonType
	if seasonType == "" {
		seasonType = "Regular Season"
	}

	reqParams := map[string]string{
		"PlayerID": params.PlayerId,
		"MeasureType": measureType,
		"PaceAdjust": paceAdjust,
		"PerMode": perMode,
		"PlusMinus": plusMinus,
		"Rank": rank,
		"Season": season,
		"SeasonType": seasonType,
	}
	
	if params.LeagueIdNullable != "" {
		reqParams["LeagueID"] = params.LeagueIdNullable
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

