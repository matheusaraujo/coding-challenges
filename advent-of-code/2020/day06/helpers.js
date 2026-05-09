export function parseGroups(puzzleInput) {
  const groups = [];
  let current = [];
  for (const line of puzzleInput) {
    if (line === "") {
      groups.push(current);
      current = [];
    } else {
      current.push(line);
    }
  }
  if (current.length) groups.push(current);
  return groups;
}
