#include "helpers.h"
#include <algorithm>
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  auto scanners = alignScanners(puzzleInput).scanners;

  int best = 0;
  for (const auto &a : scanners)
    for (const auto &b : scanners)
      best = max(best, abs(a[0] - b[0]) + abs(a[1] - b[1]) + abs(a[2] - b[2]));
  return best;
}
