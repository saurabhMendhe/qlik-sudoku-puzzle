package sudoku

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

const size = 9

// Grid represents a 9x9 Sudoku grid where 0 indicates an empty cell
type Grid [size][size]int

// Solver encapsulates the Sudoku solving logic
type Solver struct {
	grid   *Grid
	logger zap.Logger
}

func NewSolver(grid *Grid, logger zap.Logger) *Solver {
	return &Solver{
		grid:   grid,
		logger: logger,
	}
}

func (s *Solver) isValid(row, col, num int) bool {
	for c := 0; c < size; c++ {
		if s.grid[row][c] == num {
			return false
		}
	}

	for r := 0; r < size; r++ {
		if s.grid[r][col] == num {
			return false
		}
	}

	startRow, startCol := (row/3)*3, (col/3)*3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if s.grid[r][c] == num {
				return false
			}
		}
	}

	return true
}

// findEmptyCell finds the next empty cell in the grid.
func (s *Solver) findEmptyCell() (int, int) {
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if s.grid[r][c] == 0 {
				return r, c
			}
		}
	}
	return -1, -1
}

// solve uses backtracking to solve the Sudoku puzzle.
func (s *Solver) Solve() bool {
	row, col := s.findEmptyCell()
	if row == -1 && col == -1 {
		return true // Puzzle solved
	}

	for num := 1; num <= size; num++ {
		if s.isValid(row, col, num) {
			s.grid[row][col] = num
			if s.Solve() {
				return true
			}
			// Backtrack
			s.grid[row][col] = 0
		}
	}

	return false // Trigger backtracking if no solution is found
}

// PrintGrid prints the Sudoku grid.
func (s *Solver) PrintGrid() {
	var gridStr string

	// Loop through the grid
	for r := 0; r < size; r++ {
		// Loop through each column in the row
		for c := 0; c < size; c++ {
			// Add the value to the grid string
			gridStr += fmt.Sprintf("%d", s.grid[r][c])

			// Add a comma after each element except the last element in a row
			if c < size-1 {
				gridStr += ","
			}
		}
		// Add a newline after each row
		gridStr += "\n"
	}

	// Log the grid as a simple string, with a label to avoid confusion with structured logging
	s.logger.Info("Sudoku Grid:\n" + gridStr)
}

// ParseInput parses the string input into a 2D array representing the Sudoku grid.
func ParseInput(input string) (Grid, error) {
	var grid Grid
	values := strings.Split(input, ",")

	if len(values) != size*size {
		return grid, fmt.Errorf("invalid input length")
	}

	for i, v := range values {
		num := 0
		_, err := fmt.Sscanf(v, "%d", &num)
		if err != nil {
			return grid, err
		}
		row := i / size
		col := i % size
		grid[row][col] = num
	}

	return grid, nil
}

func Validate(grid Grid) error {
	for i := 0; i < size; i++ {
		if err := validateRow(grid, i); err != nil {
			return err
		}
		if err := validateColumn(grid, i); err != nil {
			return err
		}
		if err := validateBox(grid, i); err != nil {
			return err
		}
	}
	return nil
}

func validateRow(grid Grid, row int) error {
	seen := make(map[int]bool)
	for col := 0; col < size; col++ {
		num := grid[row][col]
		if num == 0 {
			continue
		}
		if num < 1 || num > size {
			return fmt.Errorf("invalid number %d at position (%d,%d)", num, row, col)
		}
		if seen[num] {
			return fmt.Errorf("duplicate %d in row %d", num, row)
		}
		seen[num] = true
	}
	return nil
}

func validateColumn(grid Grid, col int) error {
	seen := make(map[int]bool)
	for row := 0; row < size; row++ {
		num := grid[row][col]
		if num == 0 {
			continue
		}
		if seen[num] {
			return fmt.Errorf("duplicate %d in column %d", num, col)
		}
		seen[num] = true
	}
	return nil
}

func validateBox(grid Grid, boxIndex int) error {
	seen := make(map[int]bool)
	startRow := (boxIndex / 3) * 3
	startCol := (boxIndex % 3) * 3

	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			num := grid[row][col]
			if num == 0 {
				continue
			}
			if seen[num] {
				return fmt.Errorf("duplicate %d in box (%d,%d)", num, boxIndex/3, boxIndex%3)
			}
			seen[num] = true
		}
	}
	return nil
}
