#include "helpers.h"
#include <sstream>
using namespace std;

pair<vector<string>, vector<string>> parseLine(const string &line) {
  auto sep = line.find(" | ");
  stringstream ps(line.substr(0, sep));
  stringstream os(line.substr(sep + 3));
  vector<string> patterns, outputs;
  string token;
  while (ps >> token)
    patterns.push_back(token);
  while (os >> token)
    outputs.push_back(token);
  return {patterns, outputs};
}
