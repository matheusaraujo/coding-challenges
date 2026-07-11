#include "helpers.h"
using namespace std;

static long long takeBits(const string &bits, size_t &pos, int n) {
  long long value = 0;
  while (n--)
    value = value * 2 + (bits[pos++] - '0');
  return value;
}

static Packet parsePacket(const string &bits, size_t &pos) {
  Packet packet;
  packet.version = takeBits(bits, pos, 3);
  packet.typeId = takeBits(bits, pos, 3);
  packet.value = 0;

  if (packet.typeId == 4) { // literal: 5-bit groups, first bit = continue
    bool more = true;
    while (more) {
      more = takeBits(bits, pos, 1);
      packet.value = packet.value * 16 + takeBits(bits, pos, 4);
    }
    return packet;
  }

  if (takeBits(bits, pos, 1) == 0) { // 15-bit total length of sub-packets
    size_t end = takeBits(bits, pos, 15) + pos;
    while (pos < end)
      packet.subPackets.push_back(parsePacket(bits, pos));
  } else { // 11-bit number of sub-packets
    int count = takeBits(bits, pos, 11);
    while (count--)
      packet.subPackets.push_back(parsePacket(bits, pos));
  }
  return packet;
}

Packet parseTransmission(const vector<string> &puzzleInput) {
  string bits;
  for (char c : puzzleInput[0]) {
    int nibble = c <= '9' ? c - '0' : c - 'A' + 10;
    for (int b = 3; b >= 0; b--)
      bits += (nibble >> b & 1) ? '1' : '0';
  }

  size_t pos = 0;
  return parsePacket(bits, pos);
}
