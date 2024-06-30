import React from 'react';
import './ControlPanel.css';

const ControlPanel = ({ onStartGame, onSolvePuzzle, gameStarted, onClearGame, isPuzzleSolved }) => {
    return (
        <div className="control-panel">
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
            {isPuzzleSolved && (
                <div className="button-container">
                    <button className="clear-game-button" onClick={onClearGame}>
                Clear Game
                    </button>
                </div>
            )}
        </div>
    );
};



export default ControlPanel;
