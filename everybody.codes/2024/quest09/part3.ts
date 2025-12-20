import { dp } from "./helpers";

export function part3(puzzleInput: string[]): string {
  const targets = puzzleInput.map((v) => parseInt(v, 10));
  const maxTarget = Math.max(...targets);

  const dpBeetles = dp(
    maxTarget,
    [1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101],
  );

  let totalMinBeetles = 0;

  for (const T of targets) {
    let minBeetlesForT = Infinity;
    const B_start = Math.ceil((T - 100) / 2);
    const B_end = Math.floor((T + 100) / 2);
    const start = Math.max(1, B_start);
    const end = Math.min(T - 1, B_end);

    for (let B1 = start; B1 <= end; B1++) {
      const B2 = T - B1;
      const currentTotalBeetles = dpBeetles[B1] + dpBeetles[B2];
      minBeetlesForT = Math.min(minBeetlesForT, currentTotalBeetles);
    }

    totalMinBeetles += minBeetlesForT;
  }

  return totalMinBeetles.toString();
}
