#include "helpers.h"
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  auto grid = parseGrid(puzzleInput);
  int total = grid.size() * grid[0].size();
  for (int i = 1;; i++)
    if (step(grid) == total)
      return i;
}
