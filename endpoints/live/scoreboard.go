package live

import (
	"context"
	"fmt"
	"log/slog"
)

// ScoreboardGame represents a single game in the scoreboard.
type ScoreboardGame struct {
	GameID             string             `json:"gameId"`
	GameCode           string             `json:"gameCode"`
	GameStatus         int                `json:"gameStatus"`
	GameStatusText     string             `json:"gameStatusText"`
	Period             int                `json:"period"`
	GameClock          string             `json:"gameClock"`
	GameTimeUTC        string             `json:"gameTimeUTC"`
	GameEt             string             `json:"gameEt"`
	RegulationPeriods  int                `json:"regulationPeriods"`
	SeriesGameNumber   string             `json:"seriesGameNumber"`
	SeriesText         string             `json:"seriesText"`
	HomeTeam           ScoreboardTeam     `json:"homeTeam"`
	AwayTeam           ScoreboardTeam     `json:"awayTeam"`
	GameLeaders        GameLeaders        `json:"gameLeaders"`
	PbOdds             PbOdds             `json:"pbOdds"`
}

// ScoreboardTeam represents a team in the scoreboard.
type ScoreboardTeam struct {
	TeamID            int               `json:"teamId"`
	TeamName          string            `json:"teamName"`
	TeamCity          string            `json:"teamCity"`
	TeamTricode       string            `json:"teamTricode"`
	Wins              int               `json:"wins"`
	Losses            int               `json:"losses"`
	Score             int               `json:"score"`
	InBonus           *string           `json:"inBonus"`
	TimeoutsRemaining int               `json:"timeoutsRemaining"`
	Periods           []TeamPeriod      `json:"periods"`
}

// TeamPeriod represents a period score.
type TeamPeriod struct {
	Period     int    `json:"period"`
	PeriodType string `json:"periodType"`
	Score      int    `json:"score"`
}

// GameLeaders represents the game leaders.
type GameLeaders struct {
	HomeLeaders PlayerLeader `json:"homeLeaders"`
	AwayLeaders PlayerLeader `json:"awayLeaders"`
}

// PlayerLeader represents a player leader.
type PlayerLeader struct {
	PersonID     int     `json:"personId"`
	Name         string  `json:"name"`
	JerseyNum    string  `json:"jerseyNum"`
	Position     string  `json:"position"`
	TeamTricode  string  `json:"teamTricode"`
	PlayerSlug   *string `json:"playerSlug"`
	Points       int     `json:"points"`
	Rebounds     int     `json:"rebounds"`
	Assists      int     `json:"assists"`
}

// PbOdds represents odds information.
type PbOdds struct {
	Team      *string `json:"team"`
	Odds      float64 `json:"odds"`
	Suspended int     `json:"suspended"`
}

// ScoreboardData represents the scoreboard data.
type ScoreboardData struct {
	GameDate    string           `json:"gameDate"`
	LeagueID    string           `json:"leagueId"`
	LeagueName  string           `json:"leagueName"`
	Games       []ScoreboardGame `json:"games"`
}

// ScoreboardResponse represents the full scoreboard response.
type ScoreboardResponse struct {
	Meta       Meta           `json:"meta"`
	Scoreboard ScoreboardData `json:"scoreboard"`
}

// Meta represents metadata in the response.
type Meta struct {
	Version int    `json:"version"`
	Request string `json:"request"`
	Time    string `json:"time"`
	Code    int    `json:"code"`
}

// GetScoreboard fetches today's scoreboard.
func (c *Client) GetScoreboard(ctx context.Context) (*ScoreboardResponse, error) {
	c.logger.InfoContext(ctx, "Fetching NBA scoreboard")

	resp, err := c.httpClient.SendRequest(ctx, "scoreboard/todaysScoreboard_00.json", nil)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch scoreboard",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch scoreboard: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from scoreboard endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var scoreboardResp ScoreboardResponse
	if err := resp.GetJSON(&scoreboardResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal scoreboard response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched scoreboard",
		slog.Int("games_count", len(scoreboardResp.Scoreboard.Games)))

	return &scoreboardResp, nil
}

