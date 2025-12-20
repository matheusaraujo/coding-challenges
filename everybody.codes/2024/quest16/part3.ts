import { parseInput, scoreLine } from "./helpers";

export function part3(puzzleInput: string[]): string {
  const { wheels: initialWheels, jumps } = parseInput(puzzleInput);
  let wheels = initialWheels;

  wheels = wheels.map((wheel) =>
    wheel.map((w) =>
      w
        .split("")
        .filter((_, i) => i % 2 === 0)
        .join(""),
    ),
  );

  type Result = [number, number];
  const memo = new Map<string, Result>();

  function maxminScore(
    wheelOffset = 0,
    pullNumber = 0,
    pullsRemaining = 256,
  ): Result {
    const key = `${wheelOffset},${pullNumber},${pullsRemaining}`;
    if (memo.has(key)) return memo.get(key)!;

    let score = 0;
    if (pullNumber > 0) {
      let line = "";
      for (let i = 0; i < wheels.length; i++) {
        const wheel = wheels[i];
        const dist = jumps[i];
        line += wheel[(pullNumber * dist + wheelOffset) % wheel.length];
      }
      score = scoreLine(line);
    }

    let result: Result;
    if (pullsRemaining > 0) {
      const next = [-1, 0, 1].map((i) =>
        maxminScore(wheelOffset + i, pullNumber + 1, pullsRemaining - 1),
      );
      const maxVal = Math.max(...next.map((r) => r[0]));
      const minVal = Math.min(...next.map((r) => r[1]));
      result = [score + maxVal, score + minVal];
    } else {
      result = [score, score];
    }

    memo.set(key, result);
    return result;
  }

  return maxminScore().join(" ");
}
