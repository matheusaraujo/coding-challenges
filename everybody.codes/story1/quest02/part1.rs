use super::helpers::solve;
use crate::{Answer, answer};

pub fn part1(puzzle_input: &[String]) -> Answer {
    answer(solve(puzzle_input, None, true))
}
