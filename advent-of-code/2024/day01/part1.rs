use super::helpers::parse_input;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let (mut left, mut right) = parse_input(puzzle_input);
    
    left.sort();
    right.sort();
    
    let sum: i32 = left.iter()
        .zip(right.iter())
        .map(|(&l, &r)| (l - r).abs())
        .sum();

    any(sum)
}