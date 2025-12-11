import { parseInput } from "./helpers";

type CountMap = {
  [key: number]: number;
};

const part2 = (puzzleInput: string[]): number => {
  const { left, right } = parseInput(puzzleInput);

  const count: CountMap = left.reduce((acc: CountMap, value: number) => {
    acc[value] = (acc[value] || 0) + 1;
    return acc;
  }, {});

  return right.reduce((sum: number, item: number) => {
    if (count[item]) {
      sum += item * count[item];
    }
    return sum;
  }, 0);
};

export default part2;
