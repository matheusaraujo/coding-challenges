#include <string>
#include <unordered_map>
#include <vector>

void parseInput(const std::vector<std::string> &puzzleInput,
                std::vector<int> &left, std::vector<int> &right);

std::string part2(const std::vector<std::string> &puzzleInput) {
  std::vector<int> left, right;
  parseInput(puzzleInput, left, right);

  std::unordered_map<int, int> countMap;
  for (int x : left) {
    countMap[x]++;
  }

  long result = 0;
  for (int item : right) {
    result += (long)item * countMap[item];
  }

  return std::to_string(result);
}