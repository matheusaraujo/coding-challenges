#include <any>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

any part2(const vector<string> &puzzleInput) {
  string line = puzzleInput[0];
  int floor = 1;

  for (size_t i = 0; i < line.length(); i++) {
    floor += line[i] == '(' ? 1 : -1;
    if (floor == -1) {
      return i;
    }
  }

  return -1;
}