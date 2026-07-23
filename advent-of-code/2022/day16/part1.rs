use crate::helpers::{best_scores, parse_problem};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let problem = parse_problem(puzzle_input);
    let best = best_scores(&problem, 30);
    any(best.values().max().copied().unwrap_or(0))
}
