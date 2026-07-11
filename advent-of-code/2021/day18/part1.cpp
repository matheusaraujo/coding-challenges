#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  Number sum = parseNumber(puzzleInput[0]);
  for (size_t i = 1; i < puzzleInput.size(); i++)
    sum = add(sum, parseNumber(puzzleInput[i]));
  return magnitude(sum);
}
