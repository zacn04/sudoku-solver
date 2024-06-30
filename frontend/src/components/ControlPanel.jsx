import React from 'react';
import './ControlPanel.css';

const ControlPanel = ({ onStartGame, onSolvePuzzle, gameStarted }) => {
    return (
        <div className="control-panel">
            <h1 className="title">Sudoku by Zac</h1>
            {!gameStarted && (
                <div className="button-container">
                    <button className="start-game-button" onClick={onStartGame}>
                        Start Game
                    </button>
                </div>
            )}
            {gameStarted && (
                <div className="button-container">
                    <button className="solve-game-button" onClick={onSolvePuzzle}>
                        Solve Puzzle
                    </button>
                </div>
            )}
        </div>
    );
};

export default ControlPanel;
