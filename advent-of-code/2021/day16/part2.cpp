#include "helpers.h"
#include <algorithm>
#include <any>
using namespace std;

static long long evaluate(const Packet &packet) {
  vector<long long> operands;
  for (const auto &sub : packet.subPackets)
    operands.push_back(evaluate(sub));

  switch (packet.typeId) {
  case 0: { // sum
    long long total = 0;
    for (long long v : operands)
      total += v;
    return total;
  }
  case 1: { // product
    long long total = 1;
    for (long long v : operands)
      total *= v;
    return total;
  }
  case 2: // minimum
    return *min_element(operands.begin(), operands.end());
  case 3: // maximum
    return *max_element(operands.begin(), operands.end());
  case 4: // literal
    return packet.value;
  case 5: // greater than
    return operands[0] > operands[1];
  case 6: // less than
    return operands[0] < operands[1];
  default: // 7: equal to
    return operands[0] == operands[1];
  }
}

any part2(const vector<string> &puzzleInput) {
  return evaluate(parseTransmission(puzzleInput));
}
