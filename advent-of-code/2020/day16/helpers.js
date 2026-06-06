export function parse(puzzleInput) {
  const sections = puzzleInput.join("\n").split("\n\n");
  const fields = Object.fromEntries(
    sections[0].split("\n").map((line) => {
      const [name, ranges] = line.split(": ");
      const [[a, b], [c, d]] = ranges.split(" or ").map((r) =>
        r.split("-").map(Number)
      );
      return [name, (v) => (v >= a && v <= b) || (v >= c && v <= d)];
    }),
  );
  const myTicket = sections[1].split("\n")[1].split(",").map(Number);
  const nearby = sections[2].split("\n").slice(1).filter(Boolean).map((l) =>
    l.split(",").map(Number)
  );
  return { fields, myTicket, nearby };
}
