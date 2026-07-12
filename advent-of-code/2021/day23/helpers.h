#pragma once
#include <string>
#include <vector>
using namespace std;

// rooms[i] holds room i's occupants from top to bottom.
vector<string> parseRooms(const vector<string> &puzzleInput);
long long minOrganizeEnergy(const vector<string> &rooms);
