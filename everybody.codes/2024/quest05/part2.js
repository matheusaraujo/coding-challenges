import { parseInput } from "./helpers.js";

export function part2(puzzleInput) {
  const columns = parseInput(puzzleInput);
  const count = {};
  let clapIdx = 0;

  for (let r = 1; r < Infinity; r++) {
    const clapper = columns[clapIdx].shift();
    const targetColumn = columns[(clapIdx + 1) % 4];
    let moves = Math.abs((clapper % (targetColumn.length * 2)) - 1);
    if (moves > targetColumn.length) {
      moves = targetColumn.length * 2 - moves;
    }
    targetColumn.splice(moves, 0, clapper);
    clapIdx = (clapIdx + 1) % columns.length;

    const result = columns.map((x) => x[0]).join("");
    count[result] = count[result] ? count[result] + 1 : 1;
    if (count[result] === 2024) {
      return r * parseInt(result);
    }
  }

  return -1;
}
