#include "helpers.h"
using namespace std;

bool isSmall(const string &cave) { return islower(cave[0]); }

static int explore(const map<string, vector<string>> &graph,
                   map<string, int> &visits, const string &cave,
                   const RevisitPolicy &canRevisit) {
  if (cave == "end")
    return 1;

  if (isSmall(cave) && visits[cave] > 0 && !canRevisit(visits, cave))
    return 0;

  visits[cave]++;
  int paths = 0;
  for (const auto &next : graph.at(cave))
    paths += explore(graph, visits, next, canRevisit);
  visits[cave]--;

  return paths;
}

int countPaths(const vector<string> &puzzleInput,
               const RevisitPolicy &canRevisit) {
  map<string, vector<string>> graph;
  for (const auto &line : puzzleInput) {
    auto sep = line.find('-');
    string a = line.substr(0, sep), b = line.substr(sep + 1);
    graph[a].push_back(b);
    graph[b].push_back(a);
  }

  map<string, int> visits;
  return explore(graph, visits, "start", canRevisit);
}
