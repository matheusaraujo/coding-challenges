export type Tree = Record<string, string[]>;

export function parseInput(puzzleInput: string[]) {
  const tree: Tree = {};
  for (const line of puzzleInput) {
    const [node, edges] = line.split(":");
    tree[node] = edges.split(",");
  }
  return tree;
}

export function findUniquePath(tree: Tree): string[] {
  const paths: string[][] = [];
  const queue = tree["RR"].flatMap((x) =>
    tree[x] !== undefined ? [["RR", x]] : [],
  );

  while (queue.length > 0) {
    const next = queue.shift()!;
    const last = next.at(-1)!;

    const branch = tree[last].filter(
      (x: string) => x === "@" || (!next.includes(x) && tree[x] !== undefined),
    );

    branch.forEach((x: string) => {
      if (x !== "@") {
        queue.push(next.concat(x));
      } else {
        paths.push(next.concat(x));
      }
    });
  }

  return paths.find(
    (x, ix, arr) => !arr.some((y, yx) => y.length === x.length && yx !== ix),
  )!;
}
