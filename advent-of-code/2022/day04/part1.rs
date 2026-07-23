use crate::helpers::{fully_contains, parse_pair};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let count = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .filter(|line| {
            let (first, second) = parse_pair(line);
            fully_contains(first, second)
        })
        .count();
    any(count)
}
