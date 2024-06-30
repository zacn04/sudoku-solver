import React from 'react';

const PuzzleCell = ({ value, onChange }) => {
    return (
        <input
            type="text"
            value={value === 0 ? '' : value}
            onChange={(e) => onChange(Number(e.target.value) || 0)}
            className="puzzle-cell"
        />
    );
};

export default PuzzleCell;
