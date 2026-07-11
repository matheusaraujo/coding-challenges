#include "helpers.h"
#include <any>
using namespace std;

// The full cave is the original tile repeated 5x5, adding 1 risk level per
// tile step to the right or down, wrapping 9 back to 1.
static vector<vector<int>> expandGrid(const vector<vector<int>> &tile) {
  int rows = tile.size(), cols = tile[0].size();
  vector<vector<int>> grid(rows * 5, vector<int>(cols * 5));
  for (int r = 0; r < rows * 5; r++)
    for (int c = 0; c < cols * 5; c++)
      grid[r][c] = (tile[r % rows][c % cols] + r / rows + c / cols - 1) % 9 + 1;
  return grid;
}

any part2(const vector<string> &puzzleInput) {
  return lowestTotalRisk(expandGrid(parseGrid(puzzleInput)));
}
