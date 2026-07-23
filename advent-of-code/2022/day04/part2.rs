use crate::helpers::{overlaps, parse_pair};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let count = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .filter(|line| {
            let (first, second) = parse_pair(line);
            overlaps(first, second)
        })
        .count();
    any(count)
}
