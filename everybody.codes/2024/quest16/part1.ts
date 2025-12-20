import { getIthLine, parseInput } from "./helpers";

export function part1(puzzleInput: string[]): string {
  const { wheels: initialWheels, jumps } = parseInput(puzzleInput);
  let wheels = initialWheels;

  wheels = wheels.map((wheel, idx) => {
    const dist = jumps[idx];
    return wheel.map((_, i) => wheel[(i * dist) % wheel.length]);
  });

  return getIthLine(100, wheels).toString();
}
