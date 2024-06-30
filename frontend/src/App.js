import React, { useState } from 'react';
import SudokuGrid from './components/SudokuGrid';
import ControlPanel from './components/ControlPanel';
import './App.css';

function App() {
    const [puzzle, setPuzzle] = useState([]);
    const [gameStarted, setGameStarted] = useState(false);
    const [isPuzzleSolved, setisPuzzleSolved] = useState(false);

    const handleStartGame = () => {
        fetch('http://localhost:8080/api/start-game')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to start game');
                }
                return response.json();
            })
            .then(data => {
                console.log('Received puzzle data:', data); 
                setPuzzle(data.grid);
                setGameStarted(true);
                setisPuzzleSolved(false);
            })
            .catch(error => {
                console.error('Error starting game:', error);
            });
    };

    const handleSolvePuzzle = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/solve-puzzle', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ grid: puzzle }), 
            });

            if (!response.ok) {
                throw new Error('Failed to solve puzzle');
            }

            const solvedPuzzle = await response.json();
            setPuzzle(solvedPuzzle.grid); 
            setisPuzzleSolved(true);
        } catch (error) {
            console.error('Error solving puzzle:', error);
        }
    };

    const handleClearGame = () => {
        setPuzzle([]);
        setGameStarted(false);
        setisPuzzleSolved(false);
    }
    return (
        <div className="app">
            <h1 className="title">Sudoku by Zac</h1>
            <ControlPanel
                onStartGame={handleStartGame}
                onSolvePuzzle={handleSolvePuzzle}
                gameStarted={gameStarted}
                onClearGame={handleClearGame}
                isPuzzleSolved={isPuzzleSolved}
            />
            {gameStarted && puzzle.length > 0 && (
                <SudokuGrid puzzle={puzzle} setPuzzle={setPuzzle} />
            )}
        </div>
    );
}

export default App;
