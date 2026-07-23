use crate::helpers::{Step, parse_rocks, simulate_sand};
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let (rocks, max_y) = parse_rocks(puzzle_input);
    let count = simulate_sand(&rocks, |occupied, x, y| {
        if y > max_y {
            Step::Void
        } else if occupied.contains(&(x, y)) {
            Step::Blocked
        } else {
            Step::Open
        }
    });
    any(count)
}
