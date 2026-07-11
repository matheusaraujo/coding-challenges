#include "helpers.h"
#include <any>
using namespace std;

static int sumVersions(const Packet &packet) {
  int total = packet.version;
  for (const auto &sub : packet.subPackets)
    total += sumVersions(sub);
  return total;
}

any part1(const vector<string> &puzzleInput) {
  return sumVersions(parseTransmission(puzzleInput));
}
