const solve = require("./helpers");

function part3(puzzleInput) {
  return solve(puzzleInput, (arr) => Math.floor(arr.length / 2));
}

module.exports = part3;
