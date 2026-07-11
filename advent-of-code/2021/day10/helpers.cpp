#include "helpers.h"
#include <map>
using namespace std;

static const map<char, char> MATCHING{
    {')', '('}, {']', '['}, {'}', '{'}, {'>', '<'}};

char firstIllegalChar(const string &line, string &stack) {
  stack.clear();
  for (char c : line) {
    if (c == '(' || c == '[' || c == '{' || c == '<') {
      stack.push_back(c);
    } else {
      if (stack.empty() || stack.back() != MATCHING.at(c))
        return c;
      stack.pop_back();
    }
  }
  return 0;
}
