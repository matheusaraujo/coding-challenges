#include <algorithm>
#include <sstream>
#include <stdexcept>
#include <string>
#include <vector>

using namespace std;

int stringToInt(const string &str) {
  try {
    return stoi(str);
  } catch (...) {
    throw runtime_error("Could not convert to number: " + str);
  }
}

void parseInput(const vector<string> &puzzleInput, vector<int> &left,
                vector<int> &right) {
  for (const string &line : puzzleInput) {
    stringstream ss(line);
    string p1, p2;
    if (ss >> p1 >> p2) {
      left.push_back(stringToInt(p1));
      right.push_back(stringToInt(p2));
    }
  }

  sort(left.begin(), left.end());
  sort(right.begin(), right.end());
}