#include <iostream>
#include <string>
#include <vector>

std::string part2(const std::vector<std::string> &puzzleInput) {
  std::string line = puzzleInput[0];
  int floor = 0;

  for (size_t i = 0; i < line.length(); i++) {
    floor += line[i] == '(' ? 1 : -1;
    if (floor == -1) {
      return std::to_string(i + 1);
    }
  }

  return "0";
}