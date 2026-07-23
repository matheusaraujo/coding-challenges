use crate::helpers::{Step, parse_rocks, simulate_sand};
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let (rocks, max_y) = parse_rocks(puzzle_input);
    let floor_y = max_y + 2;
    let count = simulate_sand(&rocks, |occupied, x, y| {
        if y == floor_y || occupied.contains(&(x, y)) {
            Step::Blocked
        } else {
            Step::Open
        }
    });
    any(count)
}
