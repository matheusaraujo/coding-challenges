#include <iostream>
#include <string>
#include <vector>

std::string part1(const std::vector<std::string> &puzzleInput) {
  std::string line = puzzleInput[0];
  int count = 0;

  for (size_t i = 0; i < line.length(); i++) {
    count += line[i] == '(' ? 1 : -1;
  }

  return std::to_string(count);
}