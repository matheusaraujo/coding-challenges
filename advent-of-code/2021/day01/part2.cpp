#include <any>
#include <string>
#include <vector>
using namespace std;

any part2(const vector<string> &puzzleInput) {
  int count = 0;
  for (size_t i = 3; i < puzzleInput.size(); i++) {
    int prev = stoi(puzzleInput[i - 3]) + stoi(puzzleInput[i - 2]) +
               stoi(puzzleInput[i - 1]);
    int curr = stoi(puzzleInput[i - 2]) + stoi(puzzleInput[i - 1]) +
               stoi(puzzleInput[i]);
    if (curr > prev)
      count++;
  }
  return count;
}
