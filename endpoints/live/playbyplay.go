package live

import (
	"context"
	"fmt"
	"log/slog"
)

// PlayByPlayResponse represents the full play-by-play response.
type PlayByPlayResponse struct {
	Meta Meta            `json:"meta"`
	Game PlayByPlayGame  `json:"game"`
}

// PlayByPlayGame represents the game data in play-by-play.
type PlayByPlayGame struct {
	GameID  string   `json:"gameId"`
	Actions []Action `json:"actions"`
}

// Action represents a single action/play in the game.
type Action struct {
	ActionNumber         int     `json:"actionNumber"`
	Clock                string  `json:"clock"`
	TimeActual           string  `json:"timeActual"`
	Period               int     `json:"period"`
	PeriodType           string  `json:"periodType"`
	TeamID               int     `json:"teamId"`
	TeamTricode          string  `json:"teamTricode"`
	ActionType           string  `json:"actionType"`
	SubType              string  `json:"subType"`
	Descriptor           string  `json:"descriptor"`
	Qualifiers           []string `json:"qualifiers"`
	PersonID             int     `json:"personId"`
	X                    *float64 `json:"x"`
	Y                    *float64 `json:"y"`
	Possession           int     `json:"possession"`
	ScoreHome            string  `json:"scoreHome"`
	ScoreAway            string  `json:"scoreAway"`
	Edited               string  `json:"edited"`
	OrderNumber          int     `json:"orderNumber"`
	XLegacy              *int    `json:"xLegacy"`
	YLegacy              *int    `json:"yLegacy"`
	IsFieldGoal          int     `json:"isFieldGoal"`
	ShotDistance         *float64 `json:"shotDistance"`
	ShotResult           string  `json:"shotResult"`
	PointsTotal          int     `json:"pointsTotal"`
	Description          string  `json:"description"`
	PersonIdsFilter      []int   `json:"personIdsFilter"`
	AssistPersonID       int     `json:"assistPersonId"`
	AssistPlayerNameInitial string `json:"assistPlayerNameInitial"`
	AssistTotal          int     `json:"assistTotal"`
	OfficialID           int     `json:"officialId"`
	FoulDrawnPersonID    int     `json:"foulDrawnPersonId"`
	FoulPersonalTotal    int     `json:"foulPersonalTotal"`
	FoulTechnicalTotal   int     `json:"foulTechnicalTotal"`
	ShotActionNumber     int     `json:"shotActionNumber"`
	ReboundTotal         int     `json:"reboundTotal"`
	ReboundDefensiveTotal int    `json:"reboundDefensiveTotal"`
	ReboundOffensiveTotal int    `json:"reboundOffensiveTotal"`
	TurnoverTotal        int     `json:"turnoverTotal"`
	StealPersonID        int     `json:"stealPersonId"`
	Value                string  `json:"value"`
	JumpBallWonPersonID  int     `json:"jumpBallWonPersonId"`
	JumpBallLostPersonID int     `json:"jumpBallLostPersonId"`
	ShotWasBlocked       int     `json:"shotWasBlocked"`
	BlockPersonID        int     `json:"blockPersonId"`
	PlayerName           string  `json:"playerName"`
	PlayerNameI          string  `json:"playerNameI"`
}

// GetPlayByPlay fetches the play-by-play for a specific game.
func (c *Client) GetPlayByPlay(ctx context.Context, gameID string) (*PlayByPlayResponse, error) {
	c.logger.InfoContext(ctx, "Fetching NBA play-by-play",
		slog.String("game_id", gameID))

	endpoint := fmt.Sprintf("playbyplay/playbyplay_%s.json", gameID)
	resp, err := c.httpClient.SendRequest(ctx, endpoint, nil)
	if err != nil {
		c.logger.ErrorContext(ctx, "Failed to fetch play-by-play",
			slog.String("game_id", gameID),
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to fetch play-by-play: %w", err)
	}

	if !resp.IsValidJSON() {
		c.logger.ErrorContext(ctx, "Invalid JSON response from play-by-play endpoint")
		return nil, fmt.Errorf("invalid JSON response")
	}

	var pbpResp PlayByPlayResponse
	if err := resp.GetJSON(&pbpResp); err != nil {
		c.logger.ErrorContext(ctx, "Failed to unmarshal play-by-play response",
			slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.InfoContext(ctx, "Successfully fetched play-by-play",
		slog.String("game_id", gameID),
		slog.Int("actions_count", len(pbpResp.Game.Actions)))

	return &pbpResp, nil
}

