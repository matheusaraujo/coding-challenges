#include <any>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

any part1(const vector<string> &puzzleInput) {
  string line = puzzleInput[0];
  int count = 0;

  for (size_t i = 0; i < line.length(); i++) {
    count += line[i] == '(' ? 1 : -1;
  }

  return count;
}