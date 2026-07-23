use crate::helpers::find_marker;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let signal = &puzzle_input[0];
    any(find_marker(signal, 4))
}
