export function parseRules(puzzleInput) {
  const rules = {};
  for (const line of puzzleInput) {
    const [bagPart, contentsPart] = line.split(" bags contain ");
    if (contentsPart === "no other bags.") {
      rules[bagPart] = [];
    } else {
      rules[bagPart] = contentsPart
        .replace(/\./g, "")
        .split(", ")
        .map((s) => {
          const match = s.match(/^(\d+) (.+?) bags?$/);
          return { count: Number(match[1]), color: match[2] };
        });
    }
  }
  return rules;
}
