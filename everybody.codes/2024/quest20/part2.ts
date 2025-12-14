import { CHANGES, processStep } from "./helpers";

export function part2(puzzleInput: string[]): any {
  const grid = puzzleInput.map((line) => line.split(""));
  const width = grid[0].length,
    height = grid.length;
  let states: { [key: string]: number } = {};

  let starting = { x: 0, y: 0 };
  const checkpoints = [
    { x: 0, y: 0 },
    { x: 0, y: 0 },
    { x: 0, y: 0 },
  ];
  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      if (grid[y][x] === "S") starting = { x, y };
      if (grid[y][x] === "A") checkpoints[0] = { x, y };
      if (grid[y][x] === "B") checkpoints[1] = { x, y };
      if (grid[y][x] === "C") checkpoints[2] = { x, y };
      if (grid[y][x].match(/[SABC]/g)) grid[y][x] = ".";
    }
  }

  ["U", "D", "L", "R"].forEach((direction) => {
    states[`${starting.x},${starting.y},${direction},0`] = 10000;
  });

  let time = 0,
    found = false;
  while (!found) {
    time++;
    states = processStep(
      states,
      grid,
      (nextPosition, newDirection, score, gridValue, key) => {
        const newScore = score + CHANGES[gridValue];
        const parts = key.split(",");
        let checkpointIndex = parseInt(parts[3]);

        if (
          checkpointIndex < checkpoints.length &&
          nextPosition.x === checkpoints[checkpointIndex].x &&
          nextPosition.y === checkpoints[checkpointIndex].y
        )
          checkpointIndex++;

        if (
          nextPosition.x === starting.x &&
          nextPosition.y === starting.y &&
          newScore >= 10000 &&
          checkpointIndex === checkpoints.length
        ) {
          found = true;
        }

        const newKey = `${nextPosition.x},${nextPosition.y},${newDirection},${checkpointIndex}`;
        return { newKey, newScore };
      },
    );
  }

  return time.toString();
}
