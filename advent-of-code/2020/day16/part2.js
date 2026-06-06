export function part2(puzzleInput) {
  let i = 0;
  const fields = [];
  while (puzzleInput[i] !== "") {
    const [name, rest] = puzzleInput[i++].split(": ");
    const ranges = rest.split(" or ").map((r) => r.split("-").map(Number));
    fields.push({ name, ranges });
  }
  const isValid = (v, f) => f.ranges.some(([a, b]) => v >= a && v <= b);

  i += 2; // skip blank + "your ticket:"
  const myTicket = puzzleInput[i++].split(",").map(Number);

  i += 2; // skip blank + "nearby tickets:"
  const nearby = [];
  while (i < puzzleInput.length && puzzleInput[i] !== "") {
    nearby.push(puzzleInput[i++].split(",").map(Number));
  }

  const valid = nearby.filter((t) =>
    t.every((v) => fields.some((f) => isValid(v, f)))
  );

  const n = fields.length;
  const possible = Array.from(
    { length: n },
    (_, pos) =>
      new Set(
        fields
          .map((f, fi) => (valid.every((t) => isValid(t[pos], f)) ? fi : -1))
          .filter((fi) => fi !== -1),
      ),
  );

  const assigned = new Array(n).fill(-1);
  let progress = true;
  while (progress) {
    progress = false;
    for (let pos = 0; pos < n; pos++) {
      if (possible[pos].size === 1) {
        const fi = [...possible[pos]][0];
        assigned[pos] = fi;
        possible.forEach((s) => s.delete(fi));
        progress = true;
      }
    }
  }

  return fields
    .filter((f) => f.name.startsWith("departure"))
    .map((f) => myTicket[assigned.indexOf(fields.indexOf(f))])
    .reduce((a, b) => a * b, 1);
}
