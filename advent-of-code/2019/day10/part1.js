const { bestAsteroid } = require("./helpers");

function part1(puzzleInput) {
  return bestAsteroid(puzzleInput).maxVisible;
}

module.exports = part1;
