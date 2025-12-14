export interface Star {
  id: number;
  x: number;
  y: number;
}

export interface Edge {
  u: number;
  v: number;
  weight: number;
}

export function distance(
  coord1: { x: number; y: number },
  coord2: { x: number; y: number },
): number {
  return Math.abs(coord1.x - coord2.x) + Math.abs(coord1.y - coord2.y);
}

export function parseInput(puzzleInput: string[]): Star[] {
  const stars: Star[] = [];
  let starIdCounter = 1;

  const height = puzzleInput.length;

  for (let r = 0; r < height; r++) {
    const row = puzzleInput[r];
    const width = row.length;

    for (let c = 0; c < width; c++) {
      if (row[c] === "*") {
        const x = c + 1;
        const y = height - r;

        stars.push({
          id: starIdCounter++,
          x: x,
          y: y,
        });
      }
    }
  }
  return stars;
}

export class UnionFind {
  private parent: number[];
  private count: number;

  constructor(n: number) {
    this.parent = Array.from({ length: n + 1 }, (_, i) => i); // Stars are 1-indexed
    this.count = n;
  }

  find(i: number): number {
    if (this.parent[i] === i) {
      return i;
    }
    this.parent[i] = this.find(this.parent[i]);
    return this.parent[i];
  }

  union(i: number, j: number): boolean {
    const rootI = this.find(i);
    const rootJ = this.find(j);

    if (rootI !== rootJ) {
      this.parent[rootI] = rootJ;
      this.count--;
      return true;
    }
    return false;
  }

  get set_count(): number {
    return this.count;
  }

  getComponents(numStars: number): Map<number, number[]> {
    const components = new Map<number, number[]>();
    for (let i = 1; i <= numStars; i++) {
      const root = this.find(i);
      if (!components.has(root)) {
        components.set(root, []);
      }
      components.get(root)?.push(i);
    }
    return components;
  }
}

function buildMinimalTree(stars: Star[]): number {
  const numStars = stars.length;
  if (numStars <= 1) {
    return 0;
  }
  const allEdges: Edge[] = [];
  for (let i = 0; i < numStars; i++) {
    for (let j = i + 1; j < numStars; j++) {
      const starA = stars[i];
      const starB = stars[j];
      const weight = distance(starA, starB);
      allEdges.push({ u: starA.id, v: starB.id, weight: weight });
    }
  }

  allEdges.sort((a, b) => a.weight - b.weight);

  const uf = new UnionFind(numStars);
  let mstWeight = 0;
  let edgesCount = 0;

  for (const edge of allEdges) {
    if (uf.union(edge.u, edge.v)) {
      mstWeight += edge.weight;
      edgesCount++;

      if (edgesCount === numStars - 1) {
        break;
      }
    }
  }

  return mstWeight;
}

export function constellationSize(puzzleInput: string[]): number {
  const stars = parseInput(puzzleInput);
  return stars.length + buildMinimalTree(stars);
}
