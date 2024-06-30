import React, { useState } from 'react';
import SudokuGrid from './components/SudokuGrid';
import ControlPanel from './components/ControlPanel';
import './App.css';

function App() {
    const [puzzle, setPuzzle] = useState([]);
    const [gameStarted, setGameStarted] = useState(false);

    const handleStartGame = () => {
        fetch('http://localhost:8080/api/start-game')
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to start game');
                }
                return response.json();
            })
            .then(data => {
                console.log('Received puzzle data:', data); // Log received data
                setPuzzle(data.grid);
                setGameStarted(true);
            })
            .catch(error => {
                console.error('Error starting game:', error);
            });
    };

    const handleSolvePuzzle = () => {
        fetch('http://localhost:8080/api/solve-puzzle', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(puzzle), // Send the entire puzzle state
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to solve puzzle');
            }
            return response.json();
        })
        .then(data => {
            setPuzzle(data); // Update puzzle state with solved puzzle data
        })
        .catch(error => {
            console.error('Error solving puzzle:', error);
        });
    };
    
    

    return (
        <div className="app">
            <ControlPanel
                onStartGame={handleStartGame}
                onSolvePuzzle={handleSolvePuzzle}
                gameStarted={gameStarted}
            />
            {gameStarted && puzzle && <SudokuGrid puzzle={puzzle} setPuzzle={setPuzzle} />}
        </div>
    );
    
}

export default App;
