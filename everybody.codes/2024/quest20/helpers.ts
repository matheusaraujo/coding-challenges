export const DIRECTIONS: { [key: string]: { x: number; y: number } } = {
  U: { x: 0, y: -1 },
  D: { x: 0, y: 1 },
  L: { x: -1, y: 0 },
  R: { x: 1, y: 0 },
};

export const NEXT_DIRECTIONS: { [key: string]: string[] } = {
  U: ["L", "U", "R"],
  D: ["R", "D", "L"],
  L: ["U", "L", "D"],
  R: ["D", "R", "U"],
};

export const CHANGES: { [key: string]: number } = {
  ".": -1,
  "-": -2,
  "+": 1,
};

export function processStep(
  states: { [key: string]: number },
  grid: string[][],
  customProcess: (
    nextPosition: { x: number; y: number },
    newDirection: string,
    score: number,
    gridValue: string,
    key: string,
  ) => { newKey: string; newScore: number } | null,
): { [key: string]: number } {
  const next: { [key: string]: number } = {};
  const width = grid[0].length;
  const height = grid.length;

  Object.entries(states).forEach(([key, score]) => {
    const parts = key.split(",");
    const x = parseInt(parts[0]);
    const y = parseInt(parts[1]);
    const direction = parts[2];

    NEXT_DIRECTIONS[direction].forEach((newDirection) => {
      const nextPosition = {
        x: x + DIRECTIONS[newDirection].x,
        y: y + DIRECTIONS[newDirection].y,
      };

      if (
        nextPosition.x >= 0 &&
        nextPosition.x < width &&
        nextPosition.y >= 0 &&
        nextPosition.y < height &&
        grid[nextPosition.y][nextPosition.x] !== "#"
      ) {
        const gridValue = grid[nextPosition.y][nextPosition.x];

        const result = customProcess(
          nextPosition,
          newDirection,
          score,
          gridValue,
          key,
        );

        if (result) {
          const { newKey, newScore } = result;
          if (next[newKey] === undefined) next[newKey] = newScore;
          else next[newKey] = Math.max(next[newKey], newScore);
        }
      }
    });
  });

  return next;
}
