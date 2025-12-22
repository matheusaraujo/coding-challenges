import { getOpenDoors } from "./helpers.js";

export function part1(puzzleInput) {
  const passcode = puzzleInput[0];
  const target = [3, 3];
  const queue = [[0, 0, ""]];

  while (queue.length > 0) {
    const [x, y, path] = queue.shift();

    if (x === target[0] && y === target[1]) return path;

    const [up, down, left, right] = getOpenDoors(passcode, path);

    if (up && y > 0) queue.push([x, y - 1, path + "U"]);
    if (down && y < 3) queue.push([x, y + 1, path + "D"]);
    if (left && x > 0) queue.push([x - 1, y, path + "L"]);
    if (right && x < 3) queue.push([x + 1, y, path + "R"]);
  }

  return null;
}
