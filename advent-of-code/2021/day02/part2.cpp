#include <any>
#include <string>
#include <vector>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  int horiz = 0, depth = 0, aim = 0;
  for (const auto &line : puzzleInput) {
    int n = stoi(line.substr(line.rfind(' ') + 1));
    if (line[0] == 'f') {
      horiz += n;
      depth += aim * n;
    } else if (line[0] == 'd')
      aim += n;
    else
      aim -= n;
  }
  return horiz * depth;
}
