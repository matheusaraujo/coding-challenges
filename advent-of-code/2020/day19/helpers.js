export function parse(puzzleInput) {
  const sep = puzzleInput.indexOf("");
  const rules = {};
  for (const line of puzzleInput.slice(0, sep)) {
    const [id, rest] = line.split(": ");
    rules[id] = rest;
  }
  const messages = puzzleInput.slice(sep + 1);
  return { rules, messages };
}

export function buildRegex(rules, id, memo = {}) {
  if (memo[id]) return memo[id];
  const rule = rules[id];
  if (rule.startsWith('"')) return (memo[id] = rule[1]);
  const alts = rule.split(" | ").map((seq) =>
    seq.split(" ").map((ref) => buildRegex(rules, ref, memo)).join("")
  );
  return (memo[id] = alts.length === 1 ? alts[0] : `(?:${alts.join("|")})`);
}
