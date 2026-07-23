use crate::helpers::{exterior_surface_area, parse_cubes};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let cubes = parse_cubes(puzzle_input);
    any(exterior_surface_area(&cubes))
}
