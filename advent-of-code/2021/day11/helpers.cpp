#include "helpers.h"
#include <vector>
using namespace std;

vector<vector<int>> parseGrid(const vector<string> &puzzleInput) {
  vector<vector<int>> grid;
  for (const auto &line : puzzleInput) {
    vector<int> row;
    for (char c : line)
      row.push_back(c - '0');
    grid.push_back(row);
  }
  return grid;
}

int step(vector<vector<int>> &grid) {
  int rows = grid.size(), cols = grid[0].size();
  vector<pair<int, int>> toFlash;

  for (int r = 0; r < rows; r++)
    for (int c = 0; c < cols; c++)
      if (++grid[r][c] == 10)
        toFlash.push_back({r, c});

  int flashes = 0;
  while (!toFlash.empty()) {
    auto [r, c] = toFlash.back();
    toFlash.pop_back();
    flashes++;

    for (int dr = -1; dr <= 1; dr++)
      for (int dc = -1; dc <= 1; dc++) {
        int nr = r + dr, nc = c + dc;
        if (nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
            ++grid[nr][nc] == 10)
          toFlash.push_back({nr, nc});
      }
  }

  for (auto &row : grid)
    for (auto &cell : row)
      if (cell > 9)
        cell = 0;

  return flashes;
}
