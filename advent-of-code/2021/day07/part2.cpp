#include "helpers.h"
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  return minFuel(puzzleInput, [](long long d) { return d * (d + 1) / 2; });
}
