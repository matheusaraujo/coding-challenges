#include "helpers.h"
#include <map>
using namespace std;

int countOverlaps(const vector<string> &puzzleInput,
                  function<bool(int, int, int, int)> include) {
  map<pair<int, int>, int> grid;

  for (const auto &line : puzzleInput) {
    int x1, y1, x2, y2;
    sscanf(line.c_str(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2);

    if (!include(x1, y1, x2, y2))
      continue;

    int dx = (x2 > x1) - (x2 < x1);
    int dy = (y2 > y1) - (y2 < y1);
    int steps = max(abs(x2 - x1), abs(y2 - y1));
    for (int i = 0; i <= steps; i++)
      grid[{x1 + dx * i, y1 + dy * i}]++;
  }

  int count = 0;
  for (const auto &[_, v] : grid)
    if (v >= 2)
      count++;
  return count;
}
