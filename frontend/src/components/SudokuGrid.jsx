import React from 'react';
import PuzzleCell from './PuzzleCell';
import '../App.css'

const SudokuGrid = ({ puzzle, setPuzzle }) => {
    const updateCell = (row, col, value) => {
        const updatedPuzzle = [...puzzle];
        updatedPuzzle[row][col] = value;
        setPuzzle(updatedPuzzle);
    };

    if (!Array.isArray(puzzle) || puzzle.length === 0) {
        console.log(puzzle)
        return <div>Loading...</div>;
    }

    return (
        <div className="sudoku-grid">
            {puzzle.map((row, rowIndex) => (
                <div key={rowIndex} className="sudoku-row">
                    {row.map((cell, colIndex) => (
                        <PuzzleCell
                            key={colIndex}
                            value={cell}
                            onChange={(value) => updateCell(rowIndex, colIndex, value)}
                        />
                    ))}
                </div>
            ))}
        </div>
    );
};

export default SudokuGrid;
