export function parse(puzzleInput) {
  return puzzleInput.map((line) => {
    const [left, right] = line.split(" (contains ");
    return {
      ingredients: left.split(" "),
      allergens: right ? right.slice(0, -1).split(", ") : [],
    };
  });
}

export function resolve(foods) {
  // For each allergen, candidate ingredients = intersection of all foods listing it
  const candidates = new Map();
  for (const { ingredients, allergens } of foods) {
    for (const allergen of allergens) {
      const set = new Set(ingredients);
      if (!candidates.has(allergen)) {
        candidates.set(allergen, set);
      } else {
        for (const ing of candidates.get(allergen)) {
          if (!set.has(ing)) candidates.get(allergen).delete(ing);
        }
      }
    }
  }

  // Constraint propagation
  const assigned = new Map();
  let progress = true;
  while (progress) {
    progress = false;
    for (const [allergen, set] of candidates) {
      if (set.size === 1) {
        const ing = [...set][0];
        assigned.set(allergen, ing);
        for (const s of candidates.values()) s.delete(ing);
        progress = true;
      }
    }
  }
  return assigned;
}
