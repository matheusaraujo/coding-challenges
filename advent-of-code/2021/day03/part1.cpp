#include <any>
#include <string>
#include <vector>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  int bits = puzzleInput[0].size();
  int n = puzzleInput.size();
  int gamma = 0;
  for (int b = 0; b < bits; b++) {
    int ones = 0;
    for (const auto &line : puzzleInput)
      if (line[b] == '1')
        ones++;
    if (ones > n / 2)
      gamma |= (1 << (bits - 1 - b));
  }
  int epsilon = (~gamma) & ((1 << bits) - 1);
  return gamma * epsilon;
}
