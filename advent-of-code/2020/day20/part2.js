import {
  allOrientations,
  assemble,
  countMonsters,
  MONSTER_CELLS,
  parseTiles,
  stitch,
} from "./helpers.js";

export function part2(puzzleInput) {
  const image = stitch(assemble(parseTiles(puzzleInput)));
  for (const og of allOrientations(image)) {
    const monsters = countMonsters(og);
    if (monsters > 0) {
      return og.join("").split("").filter((c) => c === "#").length -
        monsters * MONSTER_CELLS.length;
    }
  }
}
