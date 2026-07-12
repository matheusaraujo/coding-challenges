#include <any>
#include <string>
#include <vector>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  vector<string> grid = puzzleInput;
  int rows = grid.size(), cols = grid[0].size();

  for (int step = 1;; step++) {
    bool moved = false;

    // East herd moves first, all at once
    vector<string> afterEast = grid;
    for (int r = 0; r < rows; r++)
      for (int c = 0; c < cols; c++)
        if (grid[r][c] == '>' && grid[r][(c + 1) % cols] == '.') {
          afterEast[r][c] = '.';
          afterEast[r][(c + 1) % cols] = '>';
          moved = true;
        }

    // Then the south herd
    vector<string> afterSouth = afterEast;
    for (int r = 0; r < rows; r++)
      for (int c = 0; c < cols; c++)
        if (afterEast[r][c] == 'v' && afterEast[(r + 1) % rows][c] == '.') {
          afterSouth[r][c] = '.';
          afterSouth[(r + 1) % rows][c] = 'v';
          moved = true;
        }

    grid = afterSouth;
    if (!moved)
      return step;
  }
}
