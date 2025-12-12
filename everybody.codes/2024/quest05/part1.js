// based on https://www.reddit.com/r/everybodycodes/comments/1gmwffb/comment/lw6f7k5/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

import { parseInput } from "./helpers.js";

export function part1(puzzleInput) {
  const columns = parseInput(puzzleInput);
  let clapIdx = 0;

  for (let r = 0; r < 10; r++) {
    const clapper = columns[clapIdx].shift();
    const targetColumn = columns[(clapIdx + 1) % 4];
    let moves = Math.abs((clapper % (targetColumn.length * 2)) - 1);
    if (moves > targetColumn.length) {
      moves = targetColumn.length * 2 - moves;
    }
    targetColumn.splice(moves, 0, clapper);
    clapIdx = (clapIdx + 1) % columns.length;
  }

  return columns.map((x) => x[0]).join("");
}
