#include "helpers.h"
#include <any>
using namespace std;

any part1(const vector<string> &puzzleInput) {
  return (int)alignScanners(puzzleInput).beacons.size();
}
