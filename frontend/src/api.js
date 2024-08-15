const API_BASE_URL = process.env.API_BASE_URL || 'https://sudoku-solver-0e1b.onrender.com';

export const startGame = async () => {
    try {
        const response = await fetch(`${API_BASE_URL}/start-game`);
        if (!response.ok) {
            throw new Error('Failed to start game');
        }
        return await response.json();
    } catch (error) {
        console.error("Error fetching start game:", error);
        throw error;
    }
};

export const solvePuzzle = async (puzzle) => {
    try {
        const response = await fetch(`${API_BASE_URL}/solve-puzzle`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ grid: puzzle }),
        });
        if (!response.ok) {
            throw new Error('Failed to solve puzzle');
        }
        return await response.json();
    } catch (error) {
        console.error("Error solving puzzle:", error);
        throw error;
    }
};
