#include <algorithm>
#include <any>
#include <queue>
#include <string>
#include <vector>
using namespace std;

static const int DR[] = {-1, 1, 0, 0};
static const int DC[] = {0, 0, -1, 1};

static int basinSize(const vector<string> &grid, vector<vector<bool>> &visited,
                     int r, int c) {
  int rows = grid.size(), cols = grid[0].size();
  int size = 0;
  queue<pair<int, int>> q;
  q.push({r, c});
  visited[r][c] = true;

  while (!q.empty()) {
    auto [cr, cc] = q.front();
    q.pop();
    size++;
    for (int d = 0; d < 4; d++) {
      int nr = cr + DR[d], nc = cc + DC[d];
      if (nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] &&
          grid[nr][nc] != '9') {
        visited[nr][nc] = true;
        q.push({nr, nc});
      }
    }
  }

  return size;
}

any part2(const vector<string> &puzzleInput) {
  int rows = puzzleInput.size(), cols = puzzleInput[0].size();
  vector<vector<bool>> visited(rows, vector<bool>(cols, false));
  vector<int> sizes;

  for (int r = 0; r < rows; r++)
    for (int c = 0; c < cols; c++)
      if (!visited[r][c] && puzzleInput[r][c] != '9')
        sizes.push_back(basinSize(puzzleInput, visited, r, c));

  sort(sizes.rbegin(), sizes.rend());
  return (long long)sizes[0] * sizes[1] * sizes[2];
}
