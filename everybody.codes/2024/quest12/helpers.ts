export interface Coord {
  x: number;
  y: number;
  char: string;
}

export interface Data {
  a: Coord;
  b: Coord;
  c: Coord;
  t: Coord[];
}

export function parseInput(puzzleInput: string[]): Data {
  const rows = puzzleInput.length;
  const cols = puzzleInput[0].length;

  let a: Coord = { x: 0, y: 0, char: "A" };
  let b: Coord = { x: 0, y: 0, char: "B" };
  let c: Coord = { x: 0, y: 0, char: "C" };
  const t: Coord[] = [];

  for (let x = 0; x < rows; x++) {
    for (let y = 0; y < cols; y++) {
      const char = puzzleInput[rows - x - 1][y];
      if (char === "A") a = { x, y, char };
      else if (char === "B") b = { x, y, char };
      else if (char === "C") c = { x, y, char };
      else if (char === "T" || char === "H") t.push({ x, y, char });
    }
  }

  return { a, b, c, t };
}

export function destroy(s: Coord, t: Coord): number {
  for (let power = 1; power < 100; power++) {
    let x = s.x;
    let y = s.y;

    for (let i = 0; i < power; i++) {
      x++;
      y++;
      if (x === t.x && y === t.y) return power;
    }

    for (let i = 0; i < power; i++) {
      y++;
      if (x === t.x && y === t.y) return power;
    }

    while (x >= 0) {
      x--;
      y++;
      if (x === t.x && y === t.y) return power;
    }
  }

  return -1;
}

export function solve(data: Data): number {
  let result = 0;
  const groups: Coord[] = [data.a, data.b, data.c];

  const targets = [...data.t].sort((t1, t2) => t2.x - t1.x);
  const tt: Record<string, number> = { T: 1, H: 2 };

  for (const t of targets) {
    for (let i = 0; i < groups.length; i++) {
      const value = destroy(groups[i], t);
      if (value !== -1) {
        result += (i + 1) * value * tt[t.char];
        break;
      }
    }
  }
  return result;
}
