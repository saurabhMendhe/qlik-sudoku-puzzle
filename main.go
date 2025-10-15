package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/saurabhMendhe/qlik-sudoku-puzzzle/config"
	"github.com/saurabhMendhe/qlik-sudoku-puzzzle/logger"
	"github.com/saurabhMendhe/qlik-sudoku-puzzzle/sudoku"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	err := run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// Set up a CLI flag for Sudoku input
	input := flag.String("input", "", "Comma-separated string representing the Sudoku grid (e.g., 5,3,0,...)")
	flag.Parse()

	// Load configuration from environment
	cfg, err := config.Load(ctx)
	if err != nil {
		log.Printf("Failed to load configuration: %v\n", err)

		return err
	}

	configData, err := cfg.GetLoggable()
	if err != nil {
		return err
	}
	log.Printf("application config : %s", configData)

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return err
	}

	// Initialize logger
	zLog, err := logger.NewLogger(cfg)
	if err != nil {
		return err
	}
	defer zLog.Sync()

	zLog.Info("Application started",
		zap.String("version", cfg.App.Version),
		zap.String("environment", cfg.App.Environment),
	)

	// Validate the input
	if *input == "" {
		zLog.Info("Please provide an input Sudoku puzzle.")
		zLog.Info("Usage: go run main.go -input \"5,3,0,0,7,0,...\"")

		return fmt.Errorf("please provide an input sudoku puzzle")
	}

	// Parse the input Sudoku puzzle
	grid, err := sudoku.ParseInput(*input)
	if err != nil {
		return err
	}

	// Validate the input
	if err := sudoku.Validate(grid); err != nil {
		return err
	}

	s := sudoku.NewSolver(&grid, *zLog)

	// Solve the puzzle
	if s.Solve() {
		zLog.Info("Solved Sudoku:")
		s.PrintGrid()
	} else {
		zLog.Info("No solution found.")
	}

	return nil
}
