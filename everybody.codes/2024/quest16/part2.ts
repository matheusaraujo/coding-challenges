import { getIthLine, GOAL, lcm, parseInput, scoreLine } from "./helpers";

export function part2(puzzleInput: string[]): string {
  const { wheels: initialWheels, jumps } = parseInput(puzzleInput);
  let wheels = initialWheels;

  wheels = wheels.map((wheel, idx) => {
    const dist = jumps[idx];
    return wheel.map((_, i) =>
      wheel[(i * dist) % wheel.length]
        .slice(0, undefined)
        .split("")
        .filter((_, j) => j % 2 === 0)
        .join(""),
    );
  });

  const loopSize = lcm(...wheels.map((w) => w.length));
  const loopCount = Math.floor(GOAL / loopSize);
  const loopRemainder = GOAL % loopSize;

  let scoreLoop = 0;
  let scoreRem = 0;

  for (let i = 1; i <= loopSize; i++) {
    scoreLoop += scoreLine(getIthLine(i, wheels));
    if (i === loopRemainder) scoreRem = scoreLoop;
  }

  return (scoreLoop * loopCount + scoreRem).toString();
}
