import { CHANGES } from "./helpers";

export function part3(puzzleInput: string[]): any {
  const grid = puzzleInput.map((line) => line.split(""));
  const width = grid[0].length,
    height = grid.length;

  let starting = { x: 0, y: 0 };
  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      if (grid[y][x] === "S") {
        starting = { x, y };
        grid[y][x] = ".";
      }
    }
  }

  let altitude = 384400,
    distance = 0,
    rightMovement = 2;
  while (altitude > 0) {
    if (rightMovement > 0) {
      starting.x++;
      rightMovement--;
    } else {
      starting.y = (starting.y + 1) % height;
      distance++;
    }

    altitude += CHANGES[grid[starting.y][starting.x]];
  }

  return distance.toString();
}
