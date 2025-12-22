import { CHANGES, processStep } from "./helpers";

export function part1(puzzleInput: string[]): any {
  const grid = puzzleInput.map((line) => line.split(""));
  const width = grid[0].length,
    height = grid.length;
  let states: { [key: string]: number } = {};

  let starting = { x: 0, y: 0 };
  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      if (grid[y][x] === "S") {
        starting = { x, y };
        grid[y][x] = ".";
      }
    }
  }

  ["U", "D", "L", "R"].forEach((direction) => {
    states[`${starting.x},${starting.y},${direction}`] = 1000;
  });

  for (let i = 0; i < 100; i++) {
    states = processStep(
      states,
      grid,
      (nextPosition, newDirection, score, gridValue) => {
        const newScore = score + CHANGES[gridValue];
        const newKey = `${nextPosition.x},${nextPosition.y},${newDirection}`;
        return { newKey, newScore };
      },
    );
  }

  return Math.max(...Object.values(states));
}
