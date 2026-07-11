#include "helpers.h"
#include <algorithm>
#include <any>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  vector<Number> numbers;
  for (const auto &line : puzzleInput)
    numbers.push_back(parseNumber(line));

  long long best = 0;
  for (size_t i = 0; i < numbers.size(); i++)
    for (size_t j = 0; j < numbers.size(); j++)
      if (i != j)
        best = max(best, magnitude(add(numbers[i], numbers[j])));
  return best;
}
