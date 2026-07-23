use crate::helpers::tail_visits;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    any(tail_visits(puzzle_input, 2))
}
