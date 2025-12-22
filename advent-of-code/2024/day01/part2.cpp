#include <any>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

void parseInput(const vector<string> &puzzleInput, vector<int> &left,
                vector<int> &right);

any part2(const vector<string> &puzzleInput) {
  vector<int> left, right;
  parseInput(puzzleInput, left, right);

  unordered_map<int, int> countMap;
  for (int x : left) {
    countMap[x]++;
  }

  long result = 0;
  for (int item : right) {
    result += (long)item * countMap[item];
  }

  return result;
}