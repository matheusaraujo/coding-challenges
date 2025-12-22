import { getOpenDoors } from "./helpers.js";

export function part2(puzzleInput) {
  const passcode = puzzleInput[0];
  const target = [3, 3];
  const stack = [[0, 0, ""]];
  let maxLength = 0;

  while (stack.length > 0) {
    const [x, y, path] = stack.pop();

    if (x === target[0] && y === target[1]) {
      maxLength = Math.max(maxLength, path.length);
      continue;
    }

    const [up, down, left, right] = getOpenDoors(passcode, path);

    if (up && y > 0) stack.push([x, y - 1, path + "U"]);
    if (down && y < 3) stack.push([x, y + 1, path + "D"]);
    if (left && x > 0) stack.push([x - 1, y, path + "L"]);
    if (right && x < 3) stack.push([x + 1, y, path + "R"]);
  }

  return maxLength;
}
