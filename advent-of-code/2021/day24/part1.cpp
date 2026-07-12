#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  string digits(14, '9');
  for (auto [i, j, delta] : digitConstraints(puzzleInput)) {
    if (delta >= 0)
      digits[i] = '9' - delta;
    else
      digits[j] = '9' + delta;
  }
  return digits;
}
