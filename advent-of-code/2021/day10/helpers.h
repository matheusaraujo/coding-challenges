#pragma once
#include <string>
#include <vector>
using namespace std;

// Returns the first illegal closing char, or 0 if the line is incomplete;
// on incomplete lines, `stack` holds the unclosed openers.
char firstIllegalChar(const string &line, string &stack);
