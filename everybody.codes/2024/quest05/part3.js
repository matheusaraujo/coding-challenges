import { parseInput } from "./helpers.js";

export function part3(puzzleInput) {
  const columns = parseInput(puzzleInput);
  const count = {};
  let clapIdx = 0;
  const cache = {};
  let topArr = [0, 0, 0, 0];

  for (let r = 1; r < Infinity; r++) {
    const state = columns.map((x) => x.join("")).join("|");
    if (cache[state]) {
      break;
    }
    cache[state] = true;
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
    const top = columns.map((x) => parseInt(x[0]));
    if (topArr[0] < top[0]) {
      topArr = top;
    } else if (topArr[0] === top[0] && topArr[1] < top[1]) {
      topArr = top;
    } else if (
      topArr[0] === top[0] &&
      topArr[1] === top[1] &&
      topArr[2] < top[2]
    ) {
      topArr = top;
    } else if (
      topArr[0] === top[0] &&
      topArr[1] === top[1] &&
      topArr[2] === top[2] &&
      topArr[3] < top[3]
    ) {
      topArr = top;
    }
  }

  return topArr.join("");
}
