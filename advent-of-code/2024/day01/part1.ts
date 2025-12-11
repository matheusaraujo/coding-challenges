import { parseInput } from "./helpers";

const part1 = (puzzleInput: string[]): number => {
  const { left, right } = parseInput(puzzleInput);
  return left.reduce(
    (sum: number, l, index) => sum + Math.abs(l - right[index]),
    0,
  );
};

export default part1;
