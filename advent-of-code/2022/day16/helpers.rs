use std::collections::{HashMap, VecDeque};

pub struct Problem {
    pub flow: Vec<i64>,
    pub dist: Vec<Vec<i64>>,
    pub dist_from_start: Vec<i64>,
}

fn bfs(start: &str, neighbors: &HashMap<String, Vec<String>>) -> HashMap<String, i64> {
    let mut dist = HashMap::new();
    dist.insert(start.to_string(), 0);
    let mut queue = VecDeque::new();
    queue.push_back(start.to_string());

    while let Some(node) = queue.pop_front() {
        let d = dist[&node];
        for next in &neighbors[&node] {
            if !dist.contains_key(next) {
                dist.insert(next.clone(), d + 1);
                queue.push_back(next.clone());
            }
        }
    }

    dist
}

pub fn parse_problem(puzzle_input: &[String]) -> Problem {
    let mut flow_rates: HashMap<String, i64> = HashMap::new();
    let mut neighbors: HashMap<String, Vec<String>> = HashMap::new();

    for line in puzzle_input.iter().filter(|line| !line.is_empty()) {
        let mut parts = line.split("; ");
        let left = parts.next().unwrap();
        let right = parts.next().unwrap();

        let name = left[6..8].to_string();
        let rate: i64 = left.rsplit('=').next().unwrap().parse().unwrap();
        let tunnels: Vec<String> = right
            .trim_start_matches("tunnels lead to valves ")
            .trim_start_matches("tunnel leads to valve ")
            .split(", ")
            .map(|s| s.to_string())
            .collect();

        flow_rates.insert(name.clone(), rate);
        neighbors.insert(name, tunnels);
    }

    let mut important: Vec<String> = flow_rates
        .iter()
        .filter(|&(_, &rate)| rate > 0)
        .map(|(name, _)| name.clone())
        .collect();
    important.sort();

    let flow: Vec<i64> = important.iter().map(|name| flow_rates[name]).collect();

    let dist_from_start: Vec<i64> = {
        let from_start = bfs("AA", &neighbors);
        important.iter().map(|name| from_start[name]).collect()
    };

    let dist: Vec<Vec<i64>> = important
        .iter()
        .map(|from| {
            let from_here = bfs(from, &neighbors);
            important.iter().map(|to| from_here[to]).collect()
        })
        .collect();

    Problem {
        flow,
        dist,
        dist_from_start,
    }
}

fn distance(problem: &Problem, from: Option<usize>, to: usize) -> i64 {
    match from {
        None => problem.dist_from_start[to],
        Some(i) => problem.dist[i][to],
    }
}

fn explore(
    problem: &Problem,
    current: Option<usize>,
    mask: u32,
    time_left: i64,
    score: i64,
    best: &mut HashMap<u32, i64>,
) {
    let entry = best.entry(mask).or_insert(i64::MIN);
    if score > *entry {
        *entry = score;
    }

    for next in 0..problem.flow.len() {
        if mask & (1 << next) != 0 {
            continue;
        }
        let travel = distance(problem, current, next);
        let new_time = time_left - travel - 1;
        if new_time <= 0 {
            continue;
        }
        let gained = problem.flow[next] * new_time;
        explore(
            problem,
            Some(next),
            mask | (1 << next),
            new_time,
            score + gained,
            best,
        );
    }
}

pub fn best_scores(problem: &Problem, time_limit: i64) -> HashMap<u32, i64> {
    let mut best = HashMap::new();
    explore(problem, None, 0, time_limit, 0, &mut best);
    best
}
