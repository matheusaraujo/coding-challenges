const part2: object = (puzzleInput: string[]) => {
  let floor = 0;
  for (let i = 0; i < puzzleInput[0].length; i += 1) {
    floor += puzzleInput[0][i] === "(" ? 1 : -1;
    if (floor === -1) return i + 1;
  }
  return 0;
};

export default part2;
