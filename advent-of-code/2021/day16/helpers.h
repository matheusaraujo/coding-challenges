#pragma once
#include <string>
#include <vector>
using namespace std;

struct Packet {
  int version;
  int typeId;
  long long value; // literal packets (typeId 4) only
  vector<Packet> subPackets;
};

Packet parseTransmission(const vector<string> &puzzleInput);
