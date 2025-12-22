use crate::{Answer, answer};
use super::helpers::eni;
use super::helpers::solve;

pub fn part1(puzzle_input: &[String]) -> Answer {
    answer(solve(puzzle_input, |n, e, m| eni(1, n, e, m)).to_string())
}
