package sudoku_test

import (
	"reflect"
	"testing"

	"github.com/saurabhMendhe/qlik-sudoku-puzzzle/sudoku"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    sudoku.Grid
		expected sudoku.Grid
		solvable bool
	}{
		{
			name: "Solvable Sudoku",
			input: sudoku.Grid{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			expected: sudoku.Grid{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
			},
			solvable: true,
		},
		{
			name: "Unsolvable Sudoku (duplicate in row)",
			input: sudoku.Grid{
				{5, 5, 0, 0, 7, 0, 0, 0, 0}, // Invalid: 5 appears twice
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9},
			},
			solvable: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := sudoku.NewSolver(&tt.input, zap.Logger{})
			result := s.Solve()

			require.Equal(t, tt.solvable, result, "Mismatch in solvability")

			if tt.solvable {
				require.True(t, reflect.DeepEqual(tt.input, tt.expected), "Solved grid does not match expected")
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		description string
		grid        sudoku.Grid
		expectError bool
	}{
		{
			name: "Valid empty grid",
			grid: sudoku.Grid{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectError: false,
			description: "Empty grid should be valid",
		},
		{
			name: "Duplicate in row",
			grid: sudoku.Grid{
				{5, 3, 5, 0, 0, 0, 0, 0, 0}, // Two 5s
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectError: true,
			description: "Duplicate in row should be invalid",
		},
		{
			name: "Duplicate in column",
			grid: sudoku.Grid{
				{5, 0, 0, 0, 0, 0, 0, 0, 0},
				{5, 0, 0, 0, 0, 0, 0, 0, 0}, // Two 5s in column
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectError: true,
			description: "Duplicate in column should be invalid",
		},
		{
			name: "Duplicate in 3x3 box",
			grid: sudoku.Grid{
				{5, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 5, 0, 0, 0, 0, 0, 0, 0}, // Two 5s in box
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectError: true,
			description: "Duplicate in 3x3 box should be invalid",
		},
		{
			name: "Invalid number (too large)",
			grid: sudoku.Grid{
				{10, 0, 0, 0, 0, 0, 0, 0, 0}, // 10 is invalid
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expectError: true,
			description: "Number greater than 9 should be invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := sudoku.Validate(tt.grid)
			require.Equal(t, tt.expectError, err != nil, tt.description)
		})
	}
}

// TestParseInput tests string to grid conversion
func TestParseInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		description string
		expectError bool
	}{
		{
			name:        "Valid 81 character string",
			input:       "5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9",
			expectError: false,
			description: "Standard format string",
		},
		{
			name:        "Too short",
			input:       "5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0",
			expectError: true,
			description: "String shorter than 81 characters",
		},
		{
			name:        "Too long",
			input:       "5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9,0,1",
			expectError: true,
			description: "String longer than 81 characters",
		},
		{
			name:        "Invalid character",
			input:       "5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9,a,b",
			expectError: true,
			description: "Contains non-numeric character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sudoku.ParseInput(tt.input)
			require.Equal(t, tt.expectError, err != nil, tt.description)
		})
	}
}

// BenchmarkSolveEasy benchmarks the solver on an easy puzzle
func BenchmarkSolveEasy(b *testing.B) {
	grid := sudoku.Grid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	for i := 0; i < b.N; i++ {
		solver := sudoku.NewSolver(&grid, zap.Logger{})
		solver.Solve()
	}
}

// BenchmarkSolveHard benchmarks the solver on a hard puzzle
func BenchmarkSolveHard(b *testing.B) {
	grid := sudoku.Grid{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 3, 0, 8, 5},
		{0, 0, 1, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 7, 0, 0, 0},
		{0, 0, 4, 0, 0, 0, 1, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 0, 0},
		{5, 0, 0, 0, 0, 0, 0, 7, 3},
		{0, 0, 2, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 9},
	}

	for i := 0; i < b.N; i++ {
		solver := sudoku.NewSolver(&grid, zap.Logger{})
		solver.Solve()
	}
}
