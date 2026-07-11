#include "helpers.h"
#include <any>
#include <map>
using namespace std;

// 4x6 glyphs used by AoC, flattened row by row.
static const map<string, char> FONT{
    {".##.#..##..######..##..#", 'A'}, {"###.#..####.#..##..####.", 'B'},
    {".##.#..##...#...#..#.##.", 'C'}, {"#####...###.#...#...####", 'E'},
    {"#####...###.#...#...#...", 'F'}, {".##.#..##...#.###..#.###", 'G'},
    {"#..##..######..##..##..#", 'H'}, {".###..#...#...#...#..###", 'I'},
    {"..##...#...#...##..#.##.", 'J'}, {"#..##.#.##..#.#.#.#.#..#", 'K'},
    {"#...#...#...#...#...####", 'L'}, {".##.#..##..##..##..#.##.", 'O'},
    {"###.#..##..####.#...#...", 'P'}, {"###.#..##..####.#.#.#..#", 'R'},
    {".####...#....##....####.", 'S'}, {"#..##..##..##..##..#.##.", 'U'},
    {"#..##..#.##...#...#...#.", 'Y'}, {"####...#..#..#..#...####", 'Z'}};

static vector<string> render(const Dots &dots) {
  int maxX = 0, maxY = 0;
  for (auto [x, y] : dots) {
    maxX = max(maxX, x);
    maxY = max(maxY, y);
  }
  vector<string> grid(maxY + 1, string(maxX + 1, '.'));
  for (auto [x, y] : dots)
    grid[y][x] = '#';
  return grid;
}

any part2(const vector<string> &puzzleInput) {
  auto [dots, folds] = parseInput(puzzleInput);
  for (const auto &fold : folds)
    dots = applyFold(dots, fold);

  auto grid = render(dots);

  string letters;
  for (size_t col = 0; col + 4 <= grid[0].size() + 1; col += 5) {
    string glyph;
    for (const auto &row : grid)
      for (size_t c = col; c < col + 4; c++)
        glyph += c < row.size() ? row[c] : '.';
    auto it = FONT.find(glyph);
    letters += it != FONT.end() ? it->second : '?';
  }

  if (letters.find('?') == string::npos)
    return letters;

  // Unknown glyph: fall back to the raw art so it can be read manually.
  string art;
  for (const auto &row : grid)
    art += "\n" + row;
  return art;
}
