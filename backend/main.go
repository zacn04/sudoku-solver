package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type SudokuPuzzle struct {
	Grid [][]int `json:"grid"`
}

func solveSudoku(board [][]int) {
	var solve func() bool
	log.Println("solving rn")
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

func isValidSudoku(board [][]int, row, col, num int) bool {

	for i := 0; i < 9; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

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

func generateSudokuPuzzle() SudokuPuzzle {
	puzzle := SudokuPuzzle{
		Grid: make([][]int, 9),
	}
	for i := range puzzle.Grid {
		puzzle.Grid[i] = make([]int, 9)
	}

	numPlacements := chooseRandNum(7, 28, 35)

	placeInitNumbers(&puzzle, numPlacements)
	return puzzle
}

func chooseRandNum(numbers ...int) int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(numbers))
	return numbers[index]
}

func placeInitNumbers(puzzle *SudokuPuzzle, numPlacements int) {
	for count := 0; count < numPlacements; count++ {
		row := rand.Intn(9)
		col := rand.Intn(9)
		value := rand.Intn(9) + 1
		puzzle.Grid[row][col] = value
		if (count+1)%7 == 0 {
			if !isValidSudoku(puzzle.Grid, row, col, value) {
				fmt.Println("Invalid puzzle state after", count+1, "placements")
				break
			}
		}
	}
}

func startGameHandler(w http.ResponseWriter, r *http.Request) {
	puzzle := generateSudokuPuzzle()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "https://zacn04.github.io/sudoku-solver")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(puzzle)
}

func solvePuzzleHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://zacn04.github.io/sudoku-solver")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	log.Printf("Incoming %s request to %s", r.Method, r.URL.Path)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	var puzzle SudokuPuzzle
	log.Println(puzzle)

	err := json.NewDecoder(r.Body).Decode(&puzzle)
	if err != nil {
		http.Error(w, "Failed to decode puzzle", http.StatusBadRequest)
		return
	}

	log.Printf("Received puzzle: %+v", puzzle)

	board := puzzle.Grid
	solveSudoku(board)
	puzzle.Grid = board

	log.Printf("Solved puzzle: %+v", puzzle)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(puzzle); err != nil {
		http.Error(w, "Failed to encode puzzle grid", http.StatusInternalServerError)
		log.Printf("Error encoding puzzle grid: %v", err)
		return
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://zacn04.github.io/sudoku-solver")
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
	http.HandleFunc("/start-game", startGameHandler)
	http.HandleFunc("/solve-puzzle", solvePuzzleHandler)
	http.Handle("/api/", corsMiddleware(http.DefaultServeMux))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP server error: %v", err)
	}
}
