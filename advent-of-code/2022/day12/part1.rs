use crate::helpers::{distances_from_end, parse_grid};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let grid = parse_grid(puzzle_input);
    let dist = distances_from_end(&grid);
    any(dist[grid.start.0][grid.start.1])
}
