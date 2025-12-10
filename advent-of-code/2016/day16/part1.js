const { fillDisk, checksum } = require("./helpers.js");

function part1(puzzleInput) {
  const state = puzzleInput[0];
  const disk = fillDisk(state, 272);
  return checksum(disk);
}

module.exports = part1;
