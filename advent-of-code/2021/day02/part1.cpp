#include <any>
#include <string>
#include <vector>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  int horiz = 0, depth = 0;
  for (const auto &line : puzzleInput) {
    int n = stoi(line.substr(line.rfind(' ') + 1));
    if (line[0] == 'f')
      horiz += n;
    else if (line[0] == 'd')
      depth += n;
    else
      depth -= n;
  }
  return horiz * depth;
}
