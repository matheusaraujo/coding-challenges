use super::helpers::solve;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    any(solve(puzzle_input, None, true))
}
