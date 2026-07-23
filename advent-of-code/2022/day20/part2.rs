use crate::helpers::{grove_coordinates, mix, parse_numbers};
use crate::{Any, any};

const DECRYPTION_KEY: i64 = 811_589_153;

pub fn part2(puzzle_input: &[String]) -> Any {
    let numbers: Vec<i64> = parse_numbers(puzzle_input)
        .iter()
        .map(|n| n * DECRYPTION_KEY)
        .collect();
    let mixed = mix(&numbers, 10);
    any(grove_coordinates(&mixed))
}
