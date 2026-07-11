#include "helpers.h"
#include <queue>
#include <tuple>
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

int lowestTotalRisk(const vector<vector<int>> &grid) {
  static const int DR[] = {-1, 1, 0, 0};
  static const int DC[] = {0, 0, -1, 1};
  int rows = grid.size(), cols = grid[0].size();

  vector<vector<int>> risk(rows, vector<int>(cols, -1));
  using Node = tuple<int, int, int>; // total risk, row, col
  priority_queue<Node, vector<Node>, greater<>> pq;
  pq.push({0, 0, 0});

  while (!pq.empty()) {
    auto [total, r, c] = pq.top();
    pq.pop();
    if (risk[r][c] != -1)
      continue;
    risk[r][c] = total;
    if (r == rows - 1 && c == cols - 1)
      return total;

    for (int d = 0; d < 4; d++) {
      int nr = r + DR[d], nc = c + DC[d];
      if (nr >= 0 && nr < rows && nc >= 0 && nc < cols && risk[nr][nc] == -1)
        pq.push({total + grid[nr][nc], nr, nc});
    }
  }
  return -1;
}
