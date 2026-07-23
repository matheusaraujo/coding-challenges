use crate::helpers::parse_pairs;
use crate::{Any, any};
use std::cmp::Ordering;

pub fn part1(puzzle_input: &[String]) -> Any {
    let total: usize = parse_pairs(puzzle_input)
        .iter()
        .enumerate()
        .filter(|(_, (left, right))| left.cmp(right) == Ordering::Less)
        .map(|(i, _)| i + 1)
        .sum();
    any(total)
}
