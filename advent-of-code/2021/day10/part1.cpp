#include "helpers.h"
#include <any>
#include <map>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  const map<char, long long> points{
      {')', 3}, {']', 57}, {'}', 1197}, {'>', 25137}};

  long long total = 0;
  string stack;
  for (const auto &line : puzzleInput) {
    char illegal = firstIllegalChar(line, stack);
    if (illegal)
      total += points.at(illegal);
  }
  return total;
}
