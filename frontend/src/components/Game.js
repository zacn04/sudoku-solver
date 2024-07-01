import React, { useState } from 'react';
import { startGame } from '../api'; 
import './Game.css'; 

function Game() {
  const [grid, setGrid] = useState([]);
  const [gameStarted, setGameStarted] = useState(false);

  const handleStartGame = async () => {
    try {
      const data = await startGame(); 
      setGrid(data.grid); 
      setGameStarted(true);
    } catch (error) {
      console.error('Error starting game:', error);
    }
  };

  return (
    <div className="game">
      <h1>Sudoku by Zac</h1>
      {!gameStarted ? (
        <button onClick={handleStartGame} className="start-button">Start Game</button>
      ) : (
        <div className="grid">
          {grid.map((row, rowIndex) => (
            <div key={rowIndex} className="row">
              {row.map((cell, cellIndex) => (
                <input
                  key={cellIndex}
                  className="cell"
                  type="number"
                  value={cell === 0 ? '' : cell} 
                  readOnly 
                />
              ))}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default Game;
