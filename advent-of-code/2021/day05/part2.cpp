#include "helpers.h"
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  return countOverlaps(puzzleInput, [](int, int, int, int) { return true; });
}
