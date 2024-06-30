package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// SudokuPuzzle represents the structure of the puzzle grid
type SudokuPuzzle struct {
	Grid [][]int `json:"grid"`
}

// solveSudoku solves the Sudoku puzzle using backtracking
func solveSudoku(board [][]int) {
	var solve func() bool

	solve = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == 0 {
					for num := 1; num <= 9; num++ {
						if isValidSudoku(board, i, j, num) {
							board[i][j] = num
							if solve() {
								return true
							}
							board[i][j] = 0
						}
					}
					return false
				}
			}
		}
		return true
	}

	solve()
}

// isValidSudoku checks if placing num at board[row][col] is valid
func isValidSudoku(board [][]int, row, col, num int) bool {
	// Check row and column
	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// startGameHandler handles the start game request
func startGameHandler(w http.ResponseWriter, r *http.Request) {
    puzzle := SudokuPuzzle{
        Grid: [][]int{
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
    }

    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    json.NewEncoder(w).Encode(puzzle)
}


// solvePuzzleHandler handles the solve puzzle request
func solvePuzzleHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming %s request to %s", r.Method, r.URL.Path)
	if r.Method == "OPTIONS" {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.WriteHeader(http.StatusOK)
        return
    }

	if r.Method == "POST" {
        var puzzle SudokuPuzzle
        err := json.NewDecoder(r.Body).Decode(&puzzle)
        if err != nil {
            http.Error(w, "Failed to decode puzzle", http.StatusBadRequest)
            return
        }

        // Solve Sudoku
        board := puzzle.Grid
        solveSudoku(board)
        puzzle.Grid = board

        // Return solved puzzle as JSON response
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if err := json.NewEncoder(w).Encode(puzzle); err != nil {
            http.Error(w, "Failed to encode puzzle grid", http.StatusInternalServerError)
            log.Printf("Error encoding puzzle grid: %v", err)
            return
        }
    }
}

// corsMiddleware sets up CORS headers for allowing requests from frontend
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "https://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}


func main() {
	http.HandleFunc("/api/start-game", startGameHandler)
	http.HandleFunc("/api/solve-puzzle", solvePuzzleHandler)
	http.Handle("/api/", corsMiddleware(http.DefaultServeMux))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
