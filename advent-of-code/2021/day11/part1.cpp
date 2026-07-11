#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  auto grid = parseGrid(puzzleInput);
  int total = 0;
  for (int i = 0; i < 100; i++)
    total += step(grid);
  return total;
}
