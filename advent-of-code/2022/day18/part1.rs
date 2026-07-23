use crate::helpers::{parse_cubes, surface_area};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let cubes = parse_cubes(puzzle_input);
    any(surface_area(&cubes))
}
