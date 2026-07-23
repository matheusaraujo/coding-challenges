use crate::helpers::{decimal_to_snafu, snafu_to_decimal};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let sum: i64 = puzzle_input
        .iter()
        .filter(|line| !line.is_empty())
        .map(|line| snafu_to_decimal(line))
        .sum();
    any(decimal_to_snafu(sum))
}
