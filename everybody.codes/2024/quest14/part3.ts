import { directions, Position, trees } from "./helpers";

export function part3(puzzleInput: string[]): string {
  const { positions, leaves, height } = trees(
    puzzleInput.map((line) => line.split(",")),
  );

  let minMurk = Infinity;
  for (let y = 0; y <= height; y++) {
    if (positions.has(`0,${y},0`)) {
      let murk = 0;
      for (let j = 0; j < leaves.length; j++)
        murk += bfs(positions, leaves[j], y);
      minMurk = Math.min(minMurk, murk);
    }
  }

  return minMurk.toString();
}

type QueueItem = {
  pos: Position;
  steps: number;
};

function bfs(positions: Set<string>, starting: Position, endY: number): number {
  const queue: QueueItem[] = [{ pos: starting, steps: 0 }];
  const visited = new Set<string>([starting.key]);

  while (queue.length !== 0) {
    const current = queue.shift();
    if (!current) break;

    const { pos, steps } = current;

    if (pos.x === 0 && pos.y === endY && pos.z === 0) {
      return steps;
    }

    Object.values(directions).forEach((direction) => {
      const npos = new Position(
        pos.x + direction.x,
        pos.y + direction.y,
        pos.z + direction.z,
      );

      if (positions.has(npos.key) && !visited.has(npos.key)) {
        visited.add(npos.key);
        queue.push({ pos: npos, steps: steps + 1 });
      }
    });
  }

  return -1;
}
