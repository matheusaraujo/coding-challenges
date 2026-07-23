use crate::helpers::{common_char, priority};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let total: u32 = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| {
            let mid = line.len() / 2;
            let (first, second) = line.split_at(mid);
            priority(common_char(&[first, second]))
        })
        .sum();
    any(total)
}
