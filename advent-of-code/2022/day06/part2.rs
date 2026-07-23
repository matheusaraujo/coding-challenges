use crate::helpers::find_marker;
use crate::{Any, any};

pub fn part2(puzzle_input: &[String]) -> Any {
    let signal = &puzzle_input[0];
    any(find_marker(signal, 14))
}
