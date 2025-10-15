# Sudoku Puzzle in Go

A high-performance Sudoku Puzzle implemented in Go using a backtracking algorithm with optimizations.

## Table of Contents
- [Algorithm Overview](#algorithm-overview)
- [Design Decisions](#design-decisions)
- [Setup](#setup)
- [Project Structure](#project-structure)
- [Usage Examples](#usage-examples)
- [Development Commands](#development-commands)
- [Testing](#testing)
- [Performance](#performance)
- [Acknowledgements](#acknowledgements)

## Algorithm Overview

#### Backtracking Algorithm

The Sudoku Puzzle implements a **recursive backtracking** algorithm, which is a depth-first search approach:

1. **Find Empty Cell**: Scan the board from top-left to bottom-right to find the first empty cell (represented by 0)
2. **Try Values**: For each empty cell, attempt to place digits 1-9
3. **Validate**: Check if the placed digit satisfies Sudoku constraints:
   - No duplicate in the same row
   - No duplicate in the same column
   - No duplicate in the same 3×3 box
4. **Recurse**: If valid, recursively attempt to solve the remaining puzzle
5. **Backtrack**: If the recursion fails (no valid solution), undo the placement and try the next digit
6. **Solution Found**: When no empty cells remain, the puzzle is solved

#### Time Complexity
- **Worst Case**: O(9^m) where m is the number of empty cells
- **Best Case**: O(m) for puzzles with unique solution paths
- **Space Complexity**: O(m) for the recursion stack

#### Why Backtracking?

I chose backtracking as my initial approach because:

1. **Simplicity**: Easy to understand and implement correctly
2. **Completeness**: Guaranteed to find a solution if one exists
3. **Memory efficient**: Only stores the grid state, minimal overhead

**Alternative approaches considered**:
- **Brute force (without backtracking)**: A pure brute-force method has to search through all possible number combinations, which is computationally expensive and time-consuming. 

## Design Decisions

#### Struct-Based Design
```go
type Solver struct {
	grid   *Grid
	logger zap.Logger
}
```

**Rationale**: Encapsulates solving state and allows tracking metrics without global variables. Supports multiple concurrent solvers.

#### Board Representation
```go
type Grid [9][9]int
```

**Rationale**: 
- **Fixed-size array** over slice for better performance (stack allocation)
- **int type** for values (0 = empty, 1-9 = digits)

#### Validation Strategy

Separated validation (`Validate()`) from solving logic:
- Pre-validate input to fail fast on invalid puzzles
- Clearer error messages for debugging

#### API Design

**Public Functions**:
- `NewSolver()`: Constructor pattern for initialization
- `Solve()`: Uses backtracking to solve the Sudoku puzzle.
- `PrintGrid()`: prints the Sudoku grid
- `Validate()`: Standalone function for Grid validation
- `ParseInput()`: ParseInput parses the string input into a 2D array representing the Sudoku grid

**Design principle**: Minimize public API surface, hide implementation details

### Error Handling

Go idiomatic error handling:
```go
if err := Validate(grid); err != nil {
    return err
}
```

## Setup

```bash
# Clone Repository
git clone https://github.com/saurabhMendhe/qlik-sudoku-puzzle.git
cd qlik-sudoku-puzzle

# Verify Go Installation
go version

# Run Tests
make test

# Build Application
make build

# Run Example
make run
```

## Project Structure

```
qlik-sudoku-puzzle/
├── main.go                    # CLI entry point
├── sudoku/
│   ├── sudoku.go              # Core solver
│   └── sudoku_test.go         # Tests
├── logger/
│   └── logger.go              # Zap logger
├── config/
│   └── config.go              # configuration for application
├── go.mod                     # Module definition
├── Makefile                   # Build automation
├── README.md                  # Full documentation
├── DESIGN.md                  # Technical design
├── PRESENTATION_GUIDE.md      # Code walkthrough guide
└── .github/
    └── workflows/
        └── test.yml          # CI/CD pipeline
```

## Usage Examples

#### Example 1: 
```bash
go run main.go -input 5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9
```

**Output:**
```
2025-10-15T11:59:55.739+0200    INFO    qlik-sudoku-puzzle/main.go:56   Application started     {"app": "qlik-sudoku-puzzle", "version": "1.0.0", "environment": "development", "version": "1.0.0", "environment": "development"}
2025-10-15T11:59:55.740+0200    INFO    qlik-sudoku-puzzle/main.go:79   Solved Sudoku:  {"app": "qlik-sudoku-puzzle", "version": "1.0.0", "environment": "development"}
2025-10-15T11:59:55.740+0200    INFO    sudoku/sudoku.go:107    Sudoku Grid:
5,3,4,6,7,8,9,1,2
6,7,2,1,9,5,3,4,8
1,9,8,3,4,2,5,6,7
8,5,9,7,6,1,4,2,3
4,2,6,8,5,3,7,9,1
7,1,3,9,2,4,8,5,6
9,6,1,5,3,7,2,8,4
2,8,7,4,1,9,6,3,5
3,4,5,2,8,6,1,7,9
        {"app": "qlik-sudoku-puzzle", "version": "1.0.0", "environment": "development"}
```

## Development Commands

```bash
# Run All Tests
make test

# Check Test Coverage
make coverage # Opens coverage.html in browser

# Run Benchmarks
make bench

# Format Code
make fmt

# Run Static Analysis
make vet

# Build and Run
make run

# Clean Build Artifacts
make clean

# Run lint
make lint

# Run security
make security
```

## Testing

#### Unit Tests
```bash
# running unit tests
go test ./sudoku -v

# running Specific Test
go test ./sudoku -v -run TestSolve


# running with race detector
go test ./sudoku -race


# coverage report
go test ./sudoku -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Performance

#### Benchmark Results

```
BenchmarkSolveEasy-8    29563350                39.50 ns/op            0 B/op          0 allocs/op
BenchmarkSolveHard-8           1        6806072254 ns/op               8 B/op          1 allocs/op
```

## Troubleshooting

### Problem: Tests failing
```bash
# Check Go version
go version  # Should be 1.19+

# Clean and rebuild
make clean
make test
```

### Problem: Build errors
```bash
# Verify module
go mod verify

# Tidy dependencies
go mod tidy

# Rebuild
go build -v
```

## Acknowledgements

- Sudoku solving algorithms research
- Go community for best practices