const { fillDisk, checksum } = require("./helpers.js");

function part2(puzzleInput) {
  const state = puzzleInput[0];
  const disk = fillDisk(state, 35651584);
  return checksum(disk);
}

module.exports = part2;
