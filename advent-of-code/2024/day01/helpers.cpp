#include <algorithm>
#include <sstream>
#include <stdexcept>
#include <string>
#include <vector>

int stringToInt(const std::string &str) {
  try {
    return std::stoi(str);
  } catch (...) {
    throw std::runtime_error("Could not convert to number: " + str);
  }
}

void parseInput(const std::vector<std::string> &puzzleInput,
                std::vector<int> &left, std::vector<int> &right) {
  for (const std::string &line : puzzleInput) {
    std::stringstream ss(line);
    std::string p1, p2;
    if (ss >> p1 >> p2) {
      left.push_back(stringToInt(p1));
      right.push_back(stringToInt(p2));
    }
  }

  std::sort(left.begin(), left.end());
  std::sort(right.begin(), right.end());
}