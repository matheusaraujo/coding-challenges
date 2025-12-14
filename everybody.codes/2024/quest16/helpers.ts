export const GOAL = 202420242024;

export type Data = {
  jumps: number[];
  wheels: string[][];
};

export function parseInput(puzzleInput: string[]): Data {
  const jumps = puzzleInput[0].split(",").map(Number);

  const wheels: string[][] = new Array(jumps.length).fill(0).map(() => []);

  for (let i = 2; i < puzzleInput.length; i++) {
    const line = puzzleInput[i];
    for (let j = 0; j < jumps.length; j++) {
      const face = line.slice(j * 4, j * 4 + 3);
      if (face.trim() !== "") wheels[j].push(face);
    }
  }

  return {
    jumps,
    wheels,
  };
}

function gcd(a: number, b: number): number {
  while (b !== 0) [a, b] = [b, a % b];
  return Math.abs(a);
}

export function lcm(...nums: number[]): number {
  return nums.reduce((a, b) => (a * b) / gcd(a, b));
}

export function getIthLine(i: number, wheels: string[][]): string {
  return wheels.map((wheel) => wheel[i % wheel.length]).join(" ");
}

export function scoreLine(line: string): number {
  const counts = new Map<string, number>();
  for (const ch of line.replace(/ /g, "")) {
    counts.set(ch, (counts.get(ch) ?? 0) + 1);
  }
  let score = 0;
  for (const v of counts.values()) {
    if (v > 2) score += v - 2;
  }
  return score;
}
