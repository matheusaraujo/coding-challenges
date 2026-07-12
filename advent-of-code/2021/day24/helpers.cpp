#include "helpers.h"
using namespace std;

vector<tuple<int, int, int>>
digitConstraints(const vector<string> &puzzleInput) {
  auto lastNumber = [&](int line) {
    return stoi(puzzleInput[line].substr(puzzleInput[line].rfind(' ')));
  };

  vector<tuple<int, int, int>> constraints;
  vector<pair<int, int>> stack; // digit index, its "add y" constant
  for (int block = 0; block < 14; block++) {
    if (lastNumber(block * 18 + 4) == 1) {
      stack.push_back({block, lastNumber(block * 18 + 15)});
    } else {
      auto [i, addY] = stack.back();
      stack.pop_back();
      constraints.push_back({i, block, addY + lastNumber(block * 18 + 5)});
    }
  }
  return constraints;
}
