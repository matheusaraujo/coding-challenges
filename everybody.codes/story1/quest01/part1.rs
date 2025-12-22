use crate::{Any, any};
use super::helpers::eni;
use super::helpers::solve;

pub fn part1(puzzle_input: &[String]) -> Any {
    any(solve(puzzle_input, |n, e, m| eni(1, n, e, m)))
}
