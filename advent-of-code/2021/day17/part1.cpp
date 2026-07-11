#include "helpers.h"
#include <algorithm>
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  auto peaks = peakHeightsOfAllHits(puzzleInput);
  return *max_element(peaks.begin(), peaks.end());
}
