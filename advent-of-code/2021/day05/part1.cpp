#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  return countOverlaps(puzzleInput, [](int x1, int y1, int x2, int y2) {
    return x1 == x2 || y1 == y2;
  });
}
