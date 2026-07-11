#pragma once
#include <functional>
#include <map>
#include <string>
#include <vector>
using namespace std;

// Decides whether an already-visited small cave may be entered again,
// given the current visit counts.
using RevisitPolicy =
    function<bool(const map<string, int> &visits, const string &cave)>;

bool isSmall(const string &cave);
int countPaths(const vector<string> &puzzleInput,
               const RevisitPolicy &canRevisit);
