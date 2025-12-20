#include <cmath>
#include <string>
#include <vector>

void parseInput(const std::vector<std::string> &puzzleInput,
                std::vector<int> &left, std::vector<int> &right);

std::string part1(const std::vector<std::string> &puzzleInput) {
  std::vector<int> left, right;
  parseInput(puzzleInput, left, right);

  long sum = 0;
  for (size_t i = 0; i < left.size(); ++i) {
    sum += std::abs(left[i] - right[i]);
  }

  return std::to_string(sum);
}