package main

import (
    "encoding/json"
    "net/http"
)

func startGame(w http.ResponseWriter, r *http.Request) {
    puzzle := generatePuzzle()
    json.NewEncoder(w).Encode(puzzle)
}

func solvePuzzle(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Puzzle [][]int `json:"puzzle"`
    }
    json.NewDecoder(r.Body).Decode(&request)
    solution := solveSudoku(request.Puzzle)
    json.NewEncoder(w).Encode(solution)
}
