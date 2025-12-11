const solve = require("./helpers");

function part1(puzzleInput) {
  return solve(puzzleInput, [
    [1, 0],
    [-1, 0],
    [0, 1],
    [0, -1],
  ]);
}

module.exports = part1;
