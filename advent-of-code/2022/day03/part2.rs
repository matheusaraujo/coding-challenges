use crate::helpers::{common_char, priority};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let lines: Vec<&str> = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| line.as_str())
        .collect();

    let total: u32 = lines
        .chunks(3)
        .map(|group| priority(common_char(group)))
        .sum();
    any(total)
}
