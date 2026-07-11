#include <any>
#include <string>
#include <vector>
using namespace std;

static const int DR[] = {-1, 1, 0, 0};
static const int DC[] = {0, 0, -1, 1};

static bool isLowPoint(const vector<string> &grid, int r, int c) {
  int rows = grid.size(), cols = grid[r].size();
  for (int d = 0; d < 4; d++) {
    int nr = r + DR[d], nc = c + DC[d];
    if (nr >= 0 && nr < rows && nc >= 0 && nc < cols &&
        grid[nr][nc] <= grid[r][c])
      return false;
  }
  return true;
}

any part1(const vector<string> &puzzleInput) {
  int total = 0;
  for (int r = 0; r < (int)puzzleInput.size(); r++)
    for (int c = 0; c < (int)puzzleInput[r].size(); c++)
      if (isLowPoint(puzzleInput, r, c))
        total += puzzleInput[r][c] - '0' + 1;
  return total;
}
