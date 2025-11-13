package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/utkonoser/nba-api-go/endpoints/live"
)

func main() {
	// Create a logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Create a live client
	client := live.NewClient(logger)

	// Get today's scoreboard
	scoreboard, err := client.GetScoreboard(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch scoreboard: %v", err)
	}

	// Print the game date
	fmt.Printf("NBA Games on %s\n", scoreboard.Scoreboard.GameDate)
	fmt.Println("==================================================")

	// Print each game
	for i, game := range scoreboard.Scoreboard.Games {
		fmt.Printf("\nGame %d:\n", i+1)
		fmt.Printf("  %s (%d-%d) @ %s (%d-%d)\n",
			game.AwayTeam.TeamCity,
			game.AwayTeam.Wins,
			game.AwayTeam.Losses,
			game.HomeTeam.TeamCity,
			game.HomeTeam.Wins,
			game.HomeTeam.Losses)
		fmt.Printf("  Score: %d - %d\n", game.AwayTeam.Score, game.HomeTeam.Score)
		fmt.Printf("  Status: %s\n", game.GameStatusText)

		if game.Period > 0 {
			fmt.Printf("  Period: %d, Clock: %s\n", game.Period, game.GameClock)
		}
	}

	fmt.Printf("\nTotal games: %d\n", len(scoreboard.Scoreboard.Games))
}
