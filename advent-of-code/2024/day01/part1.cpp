#include <any>
#include <cmath>
#include <string>
#include <vector>

using namespace std;

void parseInput(const vector<string> &puzzleInput, vector<int> &left,
                vector<int> &right);

any part1(const vector<string> &puzzleInput) {
  vector<int> left, right;
  parseInput(puzzleInput, left, right);

  long sum = 0;
  for (size_t i = 0; i < left.size(); ++i) {
    sum += abs(left[i] - right[i]);
  }

  return to_string(sum);
}