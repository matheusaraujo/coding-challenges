use crate::helpers::directory_sizes;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let total: u64 = directory_sizes(puzzle_input)
        .values()
        .filter(|&&size| size <= 100_000)
        .sum();
    any(total)
}
