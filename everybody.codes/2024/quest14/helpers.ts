export class Position {
  constructor(
    public x: number,
    public y: number,
    public z: number,
  ) {}

  get key(): string {
    return `${this.x},${this.y},${this.z}`;
  }

  move(d: { x: number; y: number; z: number }) {
    this.x += d.x;
    this.y += d.y;
    this.z += d.z;
  }
}

export const directions: {
  [key: string]: { x: number; y: number; z: number };
} = {
  U: { x: 0, y: 1, z: 0 },
  D: { x: 0, y: -1, z: 0 },
  L: { x: -1, y: 0, z: 0 },
  R: { x: 1, y: 0, z: 0 },
  F: { x: 0, y: 0, z: 1 },
  B: { x: 0, y: 0, z: -1 },
};

export function trees(trees: string[][]): {
  positions: Set<string>;
  leaves: Position[];
  height: number;
} {
  const positions = new Set<string>();
  const leaves: Position[] = [];
  let height: number = 0;

  for (const steps of trees) {
    const p: Position = new Position(0, 0, 0);
    for (const step of steps) {
      const dir = step[0];
      const distance = parseInt(step.slice(1));
      for (let i = 0; i < distance; i++) {
        p.move(directions[dir]);
        positions.add(p.key);
        height = Math.max(height, p.y);
      }
    }
    leaves.push(p);
  }

  return { positions, leaves, height };
}
