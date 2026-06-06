#include <any>
#include <string>
#include <vector>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  int count = 0;
  for (size_t i = 1; i < puzzleInput.size(); i++)
    if (stoi(puzzleInput[i]) > stoi(puzzleInput[i - 1]))
      count++;
  return count;
}
