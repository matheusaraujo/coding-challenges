#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  int count = 0;
  for (const auto &line : puzzleInput) {
    auto [_, outputs] = parseLine(line);
    for (const auto &o : outputs) {
      int len = o.size();
      if (len == 2 || len == 3 || len == 4 || len == 7)
        count++;
    }
  }
  return count;
}
