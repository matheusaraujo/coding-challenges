use crate::helpers::{grove_coordinates, mix, parse_numbers};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let numbers = parse_numbers(puzzle_input);
    let mixed = mix(&numbers, 1);
    any(grove_coordinates(&mixed))
}
