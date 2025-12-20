import { distance, Edge, parseInput, UnionFind } from "./helpers";

export function part3(puzzleInput: string[]): string {
  const stars = parseInput(puzzleInput);
  const numStars = stars.length;

  const filteredEdges: Edge[] = [];
  for (let i = 0; i < numStars; i++) {
    for (let j = i + 1; j < numStars; j++) {
      const starA = stars[i];
      const starB = stars[j];
      const weight = distance(starA, starB);
      if (weight < 6) {
        filteredEdges.push({ u: starA.id, v: starB.id, weight: weight });
      }
    }
  }

  filteredEdges.sort((a, b) => a.weight - b.weight);

  const uf = new UnionFind(numStars);
  const mstWeightByRoot = new Map<number, number>();

  for (const edge of filteredEdges) {
    const rootU = uf.find(edge.u);
    const rootV = uf.find(edge.v);

    if (rootU !== rootV) {
      const didUnion = uf.union(edge.u, edge.v);
      if (didUnion) {
        const newRoot = uf.find(edge.u);
        const weightU = mstWeightByRoot.get(rootU) || 0;
        const weightV = mstWeightByRoot.get(rootV) || 0;

        const newWeight = weightU + weightV + edge.weight;
        mstWeightByRoot.delete(rootU);
        mstWeightByRoot.delete(rootV);
        mstWeightByRoot.set(newRoot, newWeight);
      }
    }
  }

  for (let i = 1; i <= numStars; i++) {
    const root = uf.find(i);
    if (!mstWeightByRoot.has(root)) {
      mstWeightByRoot.set(root, 0);
    }
  }

  const components = uf.getComponents(numStars);
  const constellationSizes: number[] = [];

  for (const [root, starIds] of components.entries()) {
    const componentSize = starIds.length;
    const componentMstWeight = mstWeightByRoot.get(root) || 0;
    constellationSizes.push(componentSize + componentMstWeight);
  }

  constellationSizes.sort((a, b) => b - a);

  return constellationSizes
    .slice(0, 3)
    .reduce((product, size) => product * size, 1)
    .toString();
}
