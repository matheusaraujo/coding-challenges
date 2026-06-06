const MONSTER_PATTERN = [
  "                  # ",
  "#    ##    ##    ###",
  " #  #  #  #  #  #  ",
];
export const MONSTER_CELLS = MONSTER_PATTERN.flatMap((row, r) =>
  [...row].flatMap((ch, c) => (ch === "#" ? [[r, c]] : []))
);

export function parseTiles(puzzleInput) {
  const tiles = new Map();
  let id = 0, rows = [];
  for (const line of puzzleInput) {
    if (line.startsWith("Tile")) {
      id = Number(line.match(/\d+/)[0]);
      rows = [];
    } else if (line === "") {
      tiles.set(id, rows);
    } else {
      rows.push(line);
    }
  }
  if (rows.length > 0) tiles.set(id, rows);
  return tiles;
}

export function getEdges(tile) {
  const n = tile[0].length;
  return [
    tile[0],
    tile.map((r) => r[n - 1]).join(""),
    tile[tile.length - 1],
    tile.map((r) => r[0]).join(""),
  ];
}

function rev(s) {
  return [...s].reverse().join("");
}

export function normalize(edge) {
  const r = rev(edge);
  return edge < r ? edge : r;
}

export function rotate(tile) {
  const n = tile.length;
  return Array.from(
    { length: n },
    (_, i) => Array.from({ length: n }, (_, j) => tile[n - 1 - j][i]).join(""),
  );
}

export function flipH(tile) {
  return tile.map((row) => rev(row));
}

export function allOrientations(tile) {
  const result = [];
  let t = tile;
  for (let i = 0; i < 4; i++) {
    result.push(t);
    result.push(flipH(t));
    t = rotate(t);
  }
  return result;
}

export function findCorners(tiles) {
  const edgeToTiles = new Map();
  for (const [id, tile] of tiles) {
    for (const e of getEdges(tile)) {
      const ne = normalize(e);
      if (!edgeToTiles.has(ne)) edgeToTiles.set(ne, new Set());
      edgeToTiles.get(ne).add(id);
    }
  }
  const isExterior = (e) => edgeToTiles.get(normalize(e)).size === 1;
  return [...tiles.keys()].filter(
    (id) => getEdges(tiles.get(id)).filter(isExterior).length === 2,
  );
}

export function assemble(tiles) {
  const size = Math.sqrt(tiles.size);

  const edgeToTiles = new Map();
  for (const [id, tile] of tiles) {
    for (const e of getEdges(tile)) {
      const ne = normalize(e);
      if (!edgeToTiles.has(ne)) edgeToTiles.set(ne, new Set());
      edgeToTiles.get(ne).add(id);
    }
  }
  const isExterior = (e) => edgeToTiles.get(normalize(e)).size === 1;

  // Build lookup: edge string -> [{id, tile}] for all orientations
  const byTop = new Map(), byLeft = new Map();
  for (const [id, tile] of tiles) {
    for (const og of allOrientations(tile)) {
      const [top, , , left] = getEdges(og);
      (byTop.get(top) ?? byTop.set(top, []).get(top)).push({ id, tile: og });
      (byLeft.get(left) ?? byLeft.set(left, []).get(left)).push({
        id,
        tile: og,
      });
    }
  }

  // Orient top-left corner: exterior top and left
  const cornerId = findCorners(tiles)[0];
  const topLeft = allOrientations(tiles.get(cornerId)).find((og) => {
    const [top, , , left] = getEdges(og);
    return isExterior(top) && isExterior(left);
  });

  const grid = Array.from({ length: size }, () => new Array(size));
  const used = new Set([cornerId]);
  grid[0][0] = topLeft;

  for (let r = 0; r < size; r++) {
    for (let c = 0; c < size; c++) {
      if (r === 0 && c === 0) continue;
      let candidates;
      if (r === 0) {
        const req = getEdges(grid[r][c - 1])[1];
        candidates = (byLeft.get(req) || []).filter(({ id }) => !used.has(id));
      } else if (c === 0) {
        const req = getEdges(grid[r - 1][c])[2];
        candidates = (byTop.get(req) || []).filter(({ id }) => !used.has(id));
      } else {
        const reqLeft = getEdges(grid[r][c - 1])[1];
        const reqTop = getEdges(grid[r - 1][c])[2];
        candidates = (byLeft.get(reqLeft) || []).filter(
          ({ id, tile }) => !used.has(id) && getEdges(tile)[0] === reqTop,
        );
      }
      const { id, tile } = candidates[0];
      grid[r][c] = tile;
      used.add(id);
    }
  }
  return grid;
}

export function stitch(grid) {
  const image = [];
  for (const row of grid) {
    const h = row[0].length;
    for (let r = 1; r < h - 1; r++) {
      image.push(row.map((tile) => tile[r].slice(1, -1)).join(""));
    }
  }
  return image;
}

export function countMonsters(image) {
  const h = image.length, w = image[0].length;
  let count = 0;
  for (let r = 0; r <= h - 3; r++) {
    for (let c = 0; c <= w - 20; c++) {
      if (MONSTER_CELLS.every(([dr, dc]) => image[r + dr][c + dc] === "#")) {
        count++;
      }
    }
  }
  return count;
}
