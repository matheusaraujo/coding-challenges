const solve = require("./helpers");

function part2(puzzleInput) {
  return solve(puzzleInput, [
    [1, 0],
    [-1, 0],
    [0, 1],
    [0, -1],
  ]);
}

module.exports = part2;
