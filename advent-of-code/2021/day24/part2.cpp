#include "helpers.h"
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  string digits(14, '1');
  for (auto [i, j, delta] : digitConstraints(puzzleInput)) {
    if (delta >= 0)
      digits[j] = '1' + delta;
    else
      digits[i] = '1' - delta;
  }
  return digits;
}
