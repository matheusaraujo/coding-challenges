use crate::helpers::x_during_cycles;
use crate::{Any, any};

pub fn part1(puzzle_input: &[String]) -> Any {
    let history = x_during_cycles(puzzle_input);
    let total: i32 = [20, 60, 100, 140, 180, 220]
        .iter()
        .map(|&cycle| cycle as i32 * history[cycle - 1])
        .sum();
    any(total)
}
