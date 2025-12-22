import { solve } from "./helpers.js";

export function part1(puzzleInput) {
  return solve(puzzleInput, 40);
  // let row = puzzleInput[0];
  // const width = row.length;
  // const totalRows = 40;
  // let safeCount = 0;

  // for (const tile of row) if (tile === ".") safeCount++;

  // for (let r = 1; r < totalRows; r++) {
  //   let newRow = "";

  //   for (let i = 0; i < width; i++) {
  //     const left = i > 0 ? row[i - 1] : ".";
  //     const center = row[i];
  //     const right = i < width - 1 ? row[i + 1] : ".";

  //     const isTrap =
  //       (left === "^" && center === "^" && right === ".") ||
  //       (center === "^" && right === "^" && left === ".") ||
  //       (left === "^" && center === "." && right === ".") ||
  //       (right === "^" && center === "." && left === ".");

  //     newRow += isTrap ? "^" : ".";
  //   }

  //   for (const tile of newRow) if (tile === ".") safeCount++;

  //   row = newRow;
  // }

  // return safeCount;
}
